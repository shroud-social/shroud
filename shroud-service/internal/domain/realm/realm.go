package realm

type Realm struct {
	ID         uuid.UUID `json:"id"`
	URI        string    `json:"uri"`
	AllowsNSFW bool      `json:"allows_nsfw"`
	PublicKey  string    `json:"public_key"`
	Contact    string    `json:"contact"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
}
