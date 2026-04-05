package api

import (
	"services/internal/websocket"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/ws", websocket.HandleWebSocket)

	{
		v1 := router.Group("/v1")
		v1.POST("/upload-token", GetUploadToken)
	}

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
