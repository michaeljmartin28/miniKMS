package http

type CreateKeyRequest struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
}

type EncryptRequest struct {
	Plaintext      string `json:"plaintext"`
	AdditionalData string `json:"additional_data,omitempty"`
}

type DecryptRequest struct {
	Ciphertext     string `json:"ciphertext"`
	AdditionalData string `json:"additional_data,omitempty"`
}

type GenerateDataKeyRequest struct {
	AdditionalData string `json:"additional_data,omitempty"`
}

type DecryptDataKeyRequest struct {
	EncryptedDEK   string `json:"encrypted_dek"`
	AdditionalData string `json:"additional_data,omitempty"`
}
