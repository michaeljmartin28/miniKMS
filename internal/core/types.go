package core

import "time"

type CreateKeyRequest struct {
	Name      string
	Algorithm Algorithm
	// Future: KeyUsage, ProtectionLevel, Policy, Tags, etc...

}

type CreateKeyResponse struct {
	KeyID    string
	Version  int
	CreateAt time.Time
}

type EncryptRequest struct {
	KeyID          string
	Plaintext      []byte
	AdditionalData []byte
}

type EncryptResponse struct {
	Ciphertext []byte
	Version    int
	KeyID      string
	Algorithm  Algorithm
}

type DecryptRequest struct {
	KeyID          string
	Ciphertext     []byte
	AdditionalData []byte
	Version        int
}

type DecryptResponse struct {
	Plaintext []byte
}

type GenerateDataKeyRequest struct {
	KeyID          string
	AdditionalData []byte
}

type GenerateDataKeyResponse struct {
	PlaintextDEK []byte
	EncryptedDEK []byte
	Version      int
}

type DecryptDataKeyRequest struct {
	KeyID          string
	EncryptedDEK   []byte
	Version        int
	AdditionalData []byte
}

type DecryptDataKeyResponse struct {
	PlaintextDEK []byte
}

type DisableKeyRequest struct {
	KeyID string
}

type DisableKeyResponse struct {
	KeyMetadata KeyMetadata
}

type EnableKeyRequest struct {
	KeyID string
}

type EnableKeyResponse struct {
	KeyMetadata KeyMetadata
}
