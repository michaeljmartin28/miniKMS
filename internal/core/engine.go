package core

import (
	"context"
)


type EngineConfig struct {
	DefaultAlgorithm 	Algorithm	// Default encryption algorithm, e.g., "AES-256-GCM"
	MaxKeyVersions 		int			// Maximum number of key versions to retain
	DisabledDecryption 	bool		// Decrypt data with disabled keys
	
}


type Engine struct {
	Storage KeyStore
	Crypto Crypto
	Cfg EngineConfig
}

func NewEngine(storage KeyStore, crypto Crypto, cfg EngineConfig) *Engine {
	return &Engine{
		Storage: storage,
		Crypto: crypto,
		Cfg: cfg,
	}
}

// compile-time assertion to ensure Engine implements the CoreKMS interface
var _ CoreKMS = (*Engine)(nil)

func (e *Engine) CreateKey(ctx context.Context, req CreateKeyRequest) (*CreateKeyResponse, error) {
	
	return &CreateKeyResponse{}, nil
}

func (e *Engine) Encrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error) {
	// TODO: implement
	return &EncryptResponse{}, nil
}

func (e *Engine) Decrypt(ctx context.Context, req DecryptRequest) (*DecryptResponse, error) {
	// TODO: implement
	return &DecryptResponse{}, nil
}

func (e *Engine) GenerateDataKey(ctx context.Context, req GenerateDataKeyRequest) (*GenerateDataKeyResponse, error) {
	// TODO: implement
	return &GenerateDataKeyResponse{}, nil
}

func (e *Engine) DecryptDataKey(ctx context.Context, req DecryptDataKeyRequest) (*DecryptDataKeyResponse, error) {
	// TODO: implement
	return &DecryptDataKeyResponse{}, nil
}

func (e *Engine) RotateKey(ctx context.Context, keyID string) (int, error) {
	// TODO: implement
	return 0, nil
}