package httpsrv

type CreateKeyRequest struct {
	Name      string `json:"Name"`
	Algorithm string `json:"Algorithm"`
}

type EncryptRequest struct {
	Plaintext      []byte `json:"Plaintext"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type DecryptRequest struct {
	Ciphertext     []byte `json:"Ciphertext"`
	Version        uint32 `json:"Version"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type GenerateDataKeyRequest struct {
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type DecryptDataKeyRequest struct {
	EncryptedDEK   []byte `json:"EncryptedDEK"`
	Version        uint32 `json:"Version"`
	AdditionalData []byte `json:"AdditionalData,omitempty"`
}

type RotateKeyResponse struct {
	Version uint32 `json:"Version"`
}
