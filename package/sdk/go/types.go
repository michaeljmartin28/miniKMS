package kms

import (
	"time"
)

type Key struct {
	KeyID     string    `json:"KeyID"`
	Version   uint32    `json:"Version"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type CreateKeyParams struct {
	Name      string `json:"Name"`
	Algorithm string `json:"Algorithm"`
}

type CreateKeyResponse struct {
	KeyID     string    `json:"KeyID"`
	Version   uint32    `json:"Version"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type EncryptParams struct {
	Plaintext      []byte `json:"Plaintext"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type EncryptResponse struct {
	Ciphertext []byte `json:"Ciphertext"`
	Version    uint32 `json:"Version"`
	KeyID      string `json:"KeyID"`
	Algorithm  string `json:"Algorithm"`
}

type DecryptParams struct {
	Ciphertext     []byte `json:"Ciphertext"`
	Version        uint32 `json:"Version"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type DecryptResponse struct {
	Plaintext []byte `json:"Plaintext"`
}

type KeyMetadata struct {
	KeyID         string    `json:"KeyID"`
	CreatedAt     time.Time `json:"CreatedAt"`
	State         string    `json:"State"`
	Algorithm     string    `json:"Algorithm"`
	LatestVersion uint32    `json:"LatestVersion"`
}

type RotateKeyResponse struct {
	Version uint32 `json:"Version"`
}

type GenerateDataParams struct {
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type GenerateDataKeyResponse struct {
	PlaintextDEK []byte `json:"PlaintextDEK"`
	EncryptedDEK []byte `json:"EncryptedDEK"`
	Version      uint32 `json:"Version"`
}

type DecryptDataKeyParams struct {
	EncryptedDEK   []byte `json:"EncryptedDEK"`
	Version        uint32 `json:"Version"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type DecryptDataKeyResponse struct {
	PlaintextDEK []byte `json:"PlaintextDEK"`
}
