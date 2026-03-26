package kms

import "time"

type Key struct {
	KeyID     string    `json:"key_id"`
	Version   uint32    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateKeyParams struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
}

type CreateKeyResponse struct {
	KeyID     string    `json:"key_id"`
	Version   uint32    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}
