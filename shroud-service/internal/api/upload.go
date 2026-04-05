package api

import (
	"fmt"
	"net/http"
	"services/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var userUploadSecret []byte

func LoadUploadSecret(envUploadSecret string) {
	userUploadSecret = []byte(envUploadSecret)
}

type UploadRequest struct {
	UserId     string             `json:"user_id"`
	ChannelId  string             `json:"channel_id"`
	GuildId    string             `json:"guild_id"`
	FileName   string             `json:"file_name" binding:"required"`
	Size       uint32             `json:"size" binding:"required"`
	Hash       string             `json:"hash" binding:"required"`
	UploadType storage.UploadType `json:"upload_type" binding:"required"`
}

func GetUploadToken(c *gin.Context) {
	var req UploadRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: User Auth/Permission Checks
	// TODO: Deduplication
	// TODO: Size Check/Fetch User Tier Upload Size Limit
	// TODO: Commit Metadata to DB

	var path string

	switch req.UploadType {
	case storage.TypeUserAvatar:
		path = fmt.Sprintf("/avatars/%s/%s", req.UserId, req.FileName)
	case storage.TypeUserBanner:
		path = fmt.Sprintf("/banners/%s/%s", req.UserId, req.FileName)
	case storage.TypeGuildIcon:
		path = fmt.Sprintf("/guilds/%s.webp", req.GuildId)
	case storage.TypeGuildBanner:
		path = fmt.Sprintf("/guilds/%s.webp", req.GuildId)
	case storage.TypeGuildBackground:
		path = fmt.Sprintf("/guilds/%s.webp", req.GuildId)
	case storage.TypeEmote:
		emoteId, err := uuid.NewV7()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate UUID": err.Error()})
			return
		}
		path = fmt.Sprintf("/emotes/%s", emoteId.String())
	case storage.TypeSticker:
		stickerId, err := uuid.NewV7()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate UUID": err.Error()})
			return
		}
		path = fmt.Sprintf("/stickers/%s", stickerId.String())
	case storage.TypeSound:
		soundId, err := uuid.NewV7()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate UUID": err.Error()})
			return
		}
		path = fmt.Sprintf("/sounds/%s", soundId.String())
	case storage.TypeAttachment:
		path = fmt.Sprintf("/attachments/%s/%s", req.ChannelId, req.FileName)

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path": path,
	})
	tokenString, err := token.SignedString(userUploadSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate JWT": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": tokenString,
		"uri":   fmt.Sprintf("%s%s", storage.BunnyUri, path),
	})
}
