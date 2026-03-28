package httpsrv

type CreateKeyRequest struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
}

type EncryptRequest struct {
	Plaintext      []byte `json:"plaintext"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type DecryptRequest struct {
	Ciphertext     []byte `json:"ciphertext"`
	Version        uint32 `json:"version"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type GenerateDataKeyRequest struct {
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type DecryptDataKeyRequest struct {
	EncryptedDEK   []byte `json:"encryptedDEK"`
	Version        uint32 `json:"version"`
	AdditionalData []byte `json:"additionalData,omitempty"`
}

type RotateKeyResponse struct {
	Version uint32 `json:"version"`
}
