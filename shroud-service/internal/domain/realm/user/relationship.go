package user

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusNone     Status = "NONE"
	StatusPending  Status = "PENDING"
	StatusAccepted Status = "ACCEPTED"
	StatusRejected Status = "REJECTED"
	StatusBlocked  Status = "BLOCKED"
)

type Relationship struct {
	User1ID   uuid.UUID `json:"user1_id"`
	User2ID   uuid.UUID `json:"user2_id"`
	Status    Status    `json:"status"`
	Since     time.Time `json:"since"`
	ChannelID uuid.UUID `json:"channel_id"`
}
