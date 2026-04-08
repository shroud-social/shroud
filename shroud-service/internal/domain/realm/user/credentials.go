package user

import "github.com/google/uuid"

type Credentials struct {
	ID           uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`

	// MFA
	MFAEnabled  bool     `json:"mfa_enabled"`
	MFASecret   string   `json:"mfa_token"`
	BackupCodes []string `json:"backup_codes"`

	// Passkeys
	PassKeyEnabled  bool    `json:"pass_key_enabled"`
	CredentialID    []uint8 `json:"credential_id"`
	PublicKey       []uint8 `json:"public_key"`
	SignCount       uint32  `json:"sign_count"`
	AttestationType string  `json:"attestation_type"`
}
