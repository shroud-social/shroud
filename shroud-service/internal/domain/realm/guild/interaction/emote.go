package interaction

import "github.com/google/uuid"

type Emote struct {
	ID      uuid.UUID `json:"id"`
	RealmID uuid.UUID `json:"realm_id"`
	Name    string    `json:"name"`
	Image   string    `json:"image"`
}
