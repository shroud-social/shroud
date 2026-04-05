package storage

type UploadType string

const (
	TypeUserAvatar      UploadType = "USER_AVATAR"
	TypeUserBanner      UploadType = "USER_BANNER"
	TypeGuildIcon       UploadType = "GUILD_ICON"
	TypeGuildBanner     UploadType = "GUILD_BANNER"
	TypeGuildBackground UploadType = "INVITE_BACKGROUND"
	TypeEmote           UploadType = "EMOTE"
	TypeSticker         UploadType = "STICKER"
	TypeSound           UploadType = "SOUND"
	TypeAttachment      UploadType = "ATTACHMENT"
)

var UploadConfigs = map[UploadType]struct {
	MaxSize      int64
	AllowedTypes []string
	UseVolume    bool
}{
	TypeUserAvatar: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeUserBanner: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeGuildIcon: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeGuildBanner: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeGuildBackground: {
		MaxSize:      10 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeEmote: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeSticker: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
	},
	TypeSound: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"audio/webm"},
		UseVolume:    false,
	},
	TypeAttachment: {
		MaxSize:      1000 * 1024 * 1024,
		AllowedTypes: []string{"*/*"},
		UseVolume:    true,
	},
}
