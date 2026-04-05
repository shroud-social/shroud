package tier

import (
	"time"

	"github.com/google/uuid"
)

type Tier struct {
	ID         uuid.UUID `json:"id"`
	IsOfficial bool      `json:"is_official"`

	Name              string             `json:"name"`
	Badge             string             `json:"badge"`
	AnniversaryBadges []AnniversaryBadge `json:"anniversary_badges"`

	Expiry time.Time `json:"tier_expiry"`

	UploadSize          int32 `json:"upload_size"`
	ProfileAmount       int32 `json:"profile_amount"`
	AllowsGuildProfiles bool  `json:"allows_guild_profiles"`
}

type AnniversaryBadge struct {
	ElapsedTime time.Time `json:"elapsed_time"`
	Badge       string    `json:"badge"`
	Name        string    `json:"name"`
}
