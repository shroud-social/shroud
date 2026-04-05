package guild

type Tier struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Icon             string    `json:"icon"`
	RequiredSupports uint32    `json:"required_supports"`
	Expiry           time.Time `json:"expiry"`

	AllowsInviteBackground bool   `json:"invite_background"`
	AllowsBanner           bool   `json:"banner"`
	UploadSize             uint32 `json:"upload_size"`
}
