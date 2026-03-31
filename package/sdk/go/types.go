package kms

import (
	"time"
)

type KeyMetadata struct {
	KeyID         string    `json:"keyId"`
	CreatedAt     time.Time `json:"createdAt"`
	State         string    `json:"state"`
	Algorithm     string    `json:"algorithm"`
	LatestVersion uint32    `json:"latestVersion"`
}

type CreateKeyParams struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
}

type CreateKeyResponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}

type EncryptParams struct {
	Plaintext      []byte `json:"plaintext"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type EncryptResponse struct {
	Ciphertext []byte `json:"ciphertext"`
	Version    uint32 `json:"version"`
	KeyID      string `json:"keyID"`
	Algorithm  string `json:"algorithm"`
}

type DecryptParams struct {
	Ciphertext     []byte `json:"ciphertext"`
	Version        uint32 `json:"version"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type DecryptResponse struct {
	Plaintext []byte `json:"plaintext"`
}

type RotateKeyResponse struct {
	Version uint32 `json:"version"`
}

type GenerateDataParams struct {
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type GenerateDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintextDEK"`
	EncryptedDEK []byte `json:"encryptedDEK"`
	Version      uint32 `json:"version"`
}

type DecryptDataKeyParams struct {
	EncryptedDEK   []byte `json:"encryptedDEK"`
	Version        uint32 `json:"version"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type DecryptDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintextDEK"`
}

type EnableReponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}

type DisableReponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}
