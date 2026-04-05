package trust

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

type GuildEventType EventType

type RealmEventType EventType

type UserEventType EventType

type Event struct {
	ID        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Penalty   int32     `json:"penalty"`
}

type RealmEvent struct {
	Event
	Type RealmEventType `json:"type"`
}

type GuildEvent struct {
	Event
	Type GuildEventType `json:"type"`
}

type UserEvent struct {
	Event
	Type UserEventType `json:"type"`
}
