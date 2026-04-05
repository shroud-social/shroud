package interaction

import "github.com/google/uuid"

type Sound struct {
	ID      uuid.UUID `json:"id"`
	GuildID uuid.UUID `json:"guild_id"`
	Name    string    `json:"name"`
	Sound   string    `json:"sound"`
	Icon    uuid.UUID `json:"icon"`
}
