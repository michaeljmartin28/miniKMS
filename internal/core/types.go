package core

type CreateKeyRequest struct {
	Name      string    `json:"name"`
	Algorithm Algorithm `json:"algorithm"`
	// Future: KeyUsage, ProtectionLevel, Policy, Tags, etc...

}

type CreateKeyResponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}

type EncryptRequest struct {
	KeyID          string `json:"keyId"`
	Plaintext      []byte `json:"plaintext"`
	AdditionalData []byte `json:"additionalData"`
}

type EncryptResponse struct {
	Ciphertext []byte    `json:"ciphertext"`
	Version    uint32    `json:"version"`
	KeyID      string    `json:"keyId"`
	Algorithm  Algorithm `json:"algorithm"`
}

type DecryptRequest struct {
	KeyID          string `json:"keyId"`
	Ciphertext     []byte `json:"ciphertext"`
	AdditionalData []byte `json:"additionalData"`
	Version        uint32 `json:"version"`
}

type DecryptResponse struct {
	Plaintext []byte `json:"plaintext"`
}

type GenerateDataKeyRequest struct {
	KeyID          string `json:"keyId"`
	AdditionalData []byte `json:"additionalData"`
}

type GenerateDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintextDEK"`
	EncryptedDEK []byte `json:"encryptedDEK"`
	Version      uint32 `json:"version"`
}

type DecryptDataKeyRequest struct {
	KeyID          string `json:"keyId"`
	EncryptedDEK   []byte `json:"encryptedDEK"`
	Version        uint32 `json:"version"`
	AdditionalData []byte `json:"additionalData"`
}

type DecryptDataKeyResponse struct {
	PlaintextDEK []byte `json:"plaintextDEK"`
}

type DisableKeyRequest struct {
	KeyID string `json:"keyId"`
}

type DisableKeyResponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}

type EnableKeyRequest struct {
	KeyID string `json:"keyId"`
}

type EnableKeyResponse struct {
	KeyMetadata KeyMetadata `json:"keyMetadata"`
}
