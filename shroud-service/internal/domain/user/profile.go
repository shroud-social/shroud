package user

import "github.com/google/uuid"

type Profile struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	RealmID        uuid.UUID `json:"realm_id"`
	Nickname       string    `json:"nickname"`
	Pronouns       string    `json:"pronouns"`
	Avatar         string    `json:"avatar"`
	Banner         string    `json:"banner"`
	Bio            string    `json:"bio"`
	PrimaryColor   string    `json:"primary_color"`
	SecondaryColor string    `json:"secondary_color"`
	ServerBadge    uuid.UUID `json:"server_badge"`
}
