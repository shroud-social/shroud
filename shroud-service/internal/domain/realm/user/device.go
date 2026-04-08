package user

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	DeviceID   uuid.UUID `json:"device_id"`
	UserID     uuid.UUID `json:"user_id"`
	SessionID  string    `json:"session_id"`
	DeviceIP   string    `json:"device_ip"`
	LastActive time.Time `json:"last_active"`
	UserAgent  string    `json:"user_agent"`
}
