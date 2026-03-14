package core

type CreateKeyRequest struct {
	KeySpec string
}

type CreateKeyResponse struct {
	KeyID    string
	Version  int
	CreateAt int64
}

type EncryptRequest struct {
	KeyID     string
	Plaintext []byte
}

type EncryptResponse struct {
	Ciphertext []byte
	Version    int
}

type DecryptRequest struct {
	KeyID      string
	Ciphertext []byte
}

type DecryptResponse struct {
	Plaintext []byte
}

type GenerateDataKeyRequest struct {
	KeyID string
}

type GenerateDataKeyResponse struct {
	PlaintextDataKey []byte
	EncryptedDataKey []byte
	Version          int
}

type DecryptDataKeyRequest struct {
	KeyID            string
	EncryptedDataKey []byte
	Version          int
}

type DecryptDataKeyResponse struct {
	PlaintextDataKey []byte
}