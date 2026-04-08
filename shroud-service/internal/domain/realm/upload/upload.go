package upload

import (
	"services/internal/storage"

	"github.com/golang-jwt/jwt/v5"
)

type Request struct {
	UploadType storage.UploadType `json:"upload_type" binding:"required"`
	GuildId    string             `json:"guild_id"`
	ChannelId  string             `json:"channel_id"`
	FileName   string             `json:"file_name" binding:"required"`
	Size       uint32             `json:"size" binding:"required"`
	Hash       string             `json:"hash" binding:"required"`
}

type Token struct {
	Request
	UploadId string `json:"upload_id" binding:"required"`
	UserId   string `json:"user_id" binding:"required"`
	Path     string `json:"path" binding:"required"`
	jwt.RegisteredClaims
}

type Receipt struct {
	Token
}
