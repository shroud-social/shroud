package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	internalApi "services/internal/api/service/v1"
	"services/internal/comm/pubsub"
	"services/internal/domain/realm/upload"
	"services/internal/storage"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type config struct {
	UserUploadSecret []byte
	UploadUri        string
}

var conf config

func LoadUploadConf(envUploadSecret, envUploadUri string) {
	conf.UserUploadSecret = []byte(envUploadSecret)
	conf.UploadUri = envUploadUri
}

func GetUploadToken(c *gin.Context) {
	var req upload.Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Binding Error": err.Error()})
		return
	}

	// TODO: User Auth/Permission Checks
	// TODO: Size Check/Fetch User Tier Upload Size Limit

	uploadId, err := uuid.NewV7()
	userId, _ := c.Get("user_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to generate UUID": err.Error()})
		return
	}

	replacer := strings.NewReplacer(
		"{guild_id}", req.GuildId,
		"{channel_id}", req.ChannelId,
		"{user_id}", userId.(string),
		"{file_name}", fmt.Sprintf("%s-%s", uploadId, req.FileName),
	)

	uploadConfig, ok := storage.UploadConfigs[req.UploadType]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Upload type does not exist": req.UploadType})
	}

	path := replacer.Replace(uploadConfig.Path)

	claims := upload.Token{
		Request:  req,
		UserId:   userId.(string),
		UploadId: uploadId.String(),
		Path:     path,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(conf.UserUploadSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate JWT": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Authorization": tokenString,
		"uri":           fmt.Sprintf("%s%s", conf.UploadUri, path),
	})
}

type complete struct {
	Receipt string `json:"receipt" binding:"required"`
}

func ProcessUpload(c *gin.Context) {
	var com complete
	err := c.ShouldBindJSON(&com)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Missing receipt": err.Error()})
		return
	}

	token, err := jwt.ParseWithClaims(com.Receipt, &upload.Receipt{}, func(token *jwt.Token) (interface{}, error) {
		return conf.UserUploadSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, "Invalid receipt")
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(*upload.Receipt); ok {
		jsonData, err := json.Marshal(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate JSON Marshalling Data": err.Error()})
			return
		}
		err = pubsub.Connection.Publish(internalApi.SubjectUploadNew, jsonData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't publish JSON Message": err.Error()})
			return
		}
	}
}
