package guild

import (
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	ID                uuid.UUID `json:"id"`
	ChannelID         uuid.UUID `json:"channel_id"`
	GuildID           uuid.UUID `json:"guild_id"`
	RealmID           uuid.UUID `json:"realm_id"`
	Code              string    `json:"invite_code"`
	GeneratedBy       uuid.UUID `json:"generated_by"`
	GeneratedAt       time.Time `json:"generated_at"`
	ExpiresAt         time.Time `json:"expires_at"`
	InviteLimit       int       `json:"invite_limit"`
	Uses              int       `json:"uses"`
	BypassApplication bool      `json:"bypass_application"`
}
