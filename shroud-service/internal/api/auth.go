package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var userAuthSecret []byte

func LoadAuthSecret(envAuthSecret string) {
	userAuthSecret = []byte(envAuthSecret)
}

type Claims struct {
	UserId    string `json:"user_id" binding:"required"`
	SessionId string `json:"session_id" binding:"required"`
	jwt.RegisteredClaims
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Auth Token"})
			c.Abort()
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return userAuthSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("user_id", claims.UserId)
			c.Set("session_id", claims.SessionId)
			c.Next()
		}
	}

}
