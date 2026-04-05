package channel

import (
	"time"

	"github.com/google/uuid"
)

type ChannelType string

const (
	TypeText   ChannelType = "TEXT"
	TypeVoice  ChannelType = "VOICE"
	TypeForum  ChannelType = "FORUM"
	TypeThread ChannelType = "THREAD"
	TypeDM     ChannelType = "DM"
	TypeGroup  ChannelType = "GROUP"
)

type ForumDefaultLayout string

const (
	LayoutList ForumDefaultLayout = "LIST"
	LayoutGrid ForumDefaultLayout = "GRID"
)

type ForumDefaultOrder string

const (
	OrderActivity ForumDefaultOrder = "ACTIVITY"
	OrderCreation ForumDefaultOrder = "CREATION"
)

type Channel struct {
	ID       uuid.UUID `json:"channel_id"`
	ParentID uuid.UUID `json:"parent_id"`
	GuildID  uuid.UUID `json:"guild_id"`

	CreatorID    uuid.UUID `json:"creator_id"`
	CreationTime time.Time `json:"creation_time"`

	Name     string `json:"name"`
	NSFW     bool   `json:"nsfw"`
	Slowmode uint32 `json:"slowmode"`

	// Type Specific
	Type ChannelType `json:"type"`

	// Type: Text
	Topic string `json:"topic"`
	// Type: Voice
	Bitrate   uint32 `json:"bitrate"`
	UserLimit uint32 `json:"user_limit"`
	// Type: Forum
	Tags          []ForumTag         `json:"tags"`
	DefaultLayout ForumDefaultLayout `json:"default_layout"`
	DefaultOrder  ForumDefaultOrder  `json:"default_order"`
	// Type: Thread
	Locked bool `json:"locked"`
	// Type: DM

	// Type: Group
	Image string `json:"image"`
}

type ForumTag struct {
	Icon          uuid.UUID `json:"icon"`
	Name          string    `json:"name"`
	ModeratorOnly bool      `json:"moderator_only"`
}
