package user

import (
	"services/internal/domain/user/tier"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	RealmID  string    `json:"realm_id"`
	Username string    `json:"username"`

	// Tags
	Tags []string `json:"tags"`

	Tier       tier.Tier `json:"tier"`
	TierExpiry time.Time `json:"tier_expiry"`
}
