package user

import (
	"log"
	"net/http"
	apiv1 "services/internal/api/user/v1"
	"services/internal/websocket"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	initRoutesV1(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}

func initRoutesV1(router *gin.Engine) {
	{
		v1 := router.Group("/v1")
		v1.Use(apiv1.AuthMiddleware())
		v1.GET("/ws", websocket.HandleWebSocket)
		{
			users := v1.Group("/users")
			{
				user := users.Group("/:user")
				user.GET("/profile")
				user.POST("/profile")
			}
			{
				me := users.Group("/@me")
				me.GET("")
				me.GET("/profile")
				me.PATCH("")
				me.PATCH("/profile")
				{
					relationships := me.Group("/relationships")
					relationships.GET("")
					relationships.GET("/:user_2_id")
					relationships.POST("")
					relationships.DELETE("/:user_2_id")
				}
			}
		}
		{
			channels := v1.Group("/channels")
			channels.POST("")
			{
				channel := channels.Group("/:channel_id")
				channel.GET("")
				channel.PATCH("")
				channel.DELETE("")
				{
					messages := channel.Group("/messages")
					messages.GET("")
					messages.POST("")
					messages.PATCH("/:message_id")
					messages.DELETE("/:message_id")
				}
			}
		}
		{
			guilds := v1.Group("/guilds")
			guilds.POST("")
			{
				guild := guilds.Group("/:guild_id")
				guild.GET("")
				guild.PATCH("")
				guild.DELETE("")
				{
					invites := guild.Group("/invites")
					invites.GET("")
					invites.POST("")
					invites.PATCH("/:invite_id")
					invites.DELETE("/:invite_id")
				}
				{
					roles := guild.Group("/roles")
					roles.GET("")
					roles.POST("")
					roles.PATCH("/:role_id")
					roles.DELETE("/:role_id")
				}
				{
					members := guild.Group("/members")
					members.GET("")
					members.POST("")
					members.PATCH("/:member_id")
					members.DELETE("/:member_id")
				}
				{
					interactions := guild.Group("/interactions")
					{
						emotes := interactions.Group("/emotes")
						emotes.GET("")
						emotes.GET("/:emote_id")
						emotes.PATCH("/:emote_id")
						emotes.DELETE("/:emote_id")
					}
					{
						stickers := guild.Group("/stickers")
						stickers.GET("")
						stickers.GET("/:sticker_id")
						stickers.PATCH("/:sticker_id")
						stickers.DELETE("/:sticker_id")
					}
					{
						sounds := guild.Group("/sounds")
						sounds.GET("")
						sounds.GET("/:sound_id")
						sounds.PATCH("/:sound_id")
						sounds.DELETE("/:sound_id")
					}
				}
			}
		}
		{
			uploads := v1.Group("/uploads")
			uploads.POST("/request", apiv1.GetUploadToken)
			uploads.POST("/complete", apiv1.ProcessUpload)
			uploads.GET("/@me")
			uploads.GET("/:channel_id")
		}
	}

}
