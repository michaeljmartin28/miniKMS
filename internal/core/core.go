package core

import "context"

// CoreKMS defines the interface for a Key Management Service (KMS) that provides methods for
// creating keys, encrypting and decrypting data, generating and decrypting data keys,
// and rotating keys.
type CoreKMS interface {
	CreateKey(ctx context.Context, req CreateKeyRequest) (*CreateKeyResponse, error)
	Encrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error)
	Decrypt(ctx context.Context, req DecryptRequest) (*DecryptResponse, error)
	GenerateDataKey(ctx context.Context, req GenerateDataKeyRequest) (*GenerateDataKeyResponse, error)
	DecryptDataKey(ctx context.Context, req DecryptDataKeyRequest) (*DecryptDataKeyResponse, error)
	RotateKey(ctx context.Context, keyID string) (uint32, error)
}
