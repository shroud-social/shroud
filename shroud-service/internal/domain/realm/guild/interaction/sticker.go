package interaction

import "github.com/google/uuid"

type Sticker struct {
	ID      uuid.UUID `json:"id"`
	GuildID uuid.UUID `json:"guild_id"`
	Name    string    `json:"name"`
	Image   string    `json:"image"`
}
