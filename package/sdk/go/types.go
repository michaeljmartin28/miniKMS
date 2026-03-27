package kms

import (
	"time"
)

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

type EncryptParams struct {
	Plaintext      []byte `json:"plaintext"`
	AdditionalData string `json:"additional_data,omitempty"`
}

type EncryptResponse struct {
	Ciphertext []byte `json:"ciphertext"`
	Version    uint32 `json:"version"`
	KeyID      string `json:"keyID"`
	Algorithm  string `json:"algorithm"`
}

type DecryptParams struct {
	Ciphertext     string `json:"ciphertext"`
	Version        uint32 `json:"version"`
	AdditionalData string `json:"additional_data,omitempty"`
}

type DecryptResponse struct {
	Plaintext []byte `json:"plaintext"`
}

type KeyMetadata struct {
	KeyID         string    `json:"key_id"`
	CreatedAt     time.Time `json:"created_at"`
	State         string    `json:"state"`
	Algorithm     string    `json:"algorithm"`
	LatestVersion uint32    `json:"latest_version"`
}

type RotateKeyResponse struct {
	Version uint32 `json:"version"`
}

type GenerateDataParams struct {
	AdditionalData string `json:"additional_data,omitempty"`
}

type GenerateDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintext_dek"`
	EncryptedDEK []byte `json:"encrypted_dek"`
	Version      uint32 `json:"version"`
}

type DecryptDataKeyParams struct {
	EncryptedDEK   string `json:"encrypted_dek"`
	Version        uint32 `json:"version"`
	AdditionalData string `json:"additional_data,omitempty"`
}

type DecryptDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintext_dek"`
}
