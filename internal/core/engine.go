package core

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EngineConfig struct {
	DefaultAlgorithm     Algorithm // Default encryption algorithm, e.g., "AES-256-GCM"
	MaxKeyVersions       int       // Maximum number of key versions to retain
	AllowDecryptDisabled bool      // Decrypt data with disabled keys
}

type Engine struct {
	Storage KeyStore
	Crypto  Crypto
	Cfg     EngineConfig
}

func NewEngine(storage KeyStore, crypto Crypto, cfg EngineConfig) *Engine {
	return &Engine{
		Storage: storage,
		Crypto:  crypto,
		Cfg:     cfg,
	}
}

// compile-time assertion to ensure Engine implements the CoreKMS interface
var _ CoreKMS = (*Engine)(nil)

func (e *Engine) CreateKey(ctx context.Context, req CreateKeyRequest) (*CreateKeyResponse, error) {

	keyBytes, err := e.Crypto.GenerateKey(req.Algorithm)
	version := 1

	keyVersion := KeyVersion{
		Version:   version,
		CreatedAt: time.Now(),
		Material:  keyBytes,
	}

	meta := KeyMetadata{
		KeyID:         uuid.New().String(),
		CreatedAt:     keyVersion.CreatedAt,
		Algorithm:     req.Algorithm,
		State:         Disabled,
		LatestVersion: version,
	}

	err = e.Storage.SaveKey(meta)
	if err != nil {
		return nil, err
	}

	err = e.Storage.SaveVersion(meta.KeyID, keyVersion)
	if err != nil {
		return nil, err
	}

	response := CreateKeyResponse{
		KeyID:    meta.KeyID,
		Version:  version,
		CreateAt: keyVersion.CreatedAt,
	}

	return &response, nil
}

func (e *Engine) Encrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error) {

	if req.Plaintext == nil || len(req.Plaintext) == 0 {
		return nil, fmt.Errorf("the encrypt request did not include any plaintext to encrypt")
	}

	keyMetadata, err := e.Storage.GetKey(req.KeyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State == Disabled {
		return nil, fmt.Errorf("key is disabled, cannot complete request")
	}

	keyVersion, err := e.Storage.GetVersion(keyMetadata.KeyID, keyMetadata.LatestVersion)
	if err != nil {
		return nil, err
	}

	ciphertext, err := e.Crypto.Encrypt(keyMetadata.Algorithm, keyVersion.Material, req.Plaintext, req.AdditionalData)

	if err != nil {
		return nil, err
	}

	response := &EncryptResponse{
		Ciphertext: ciphertext,
		Version:    keyVersion.Version,
		KeyID:      keyMetadata.KeyID,
		Algorithm:  keyMetadata.Algorithm,
	}

	return response, nil
}

func (e *Engine) Decrypt(ctx context.Context, req DecryptRequest) (*DecryptResponse, error) {

	if req.Ciphertext == nil || len(req.Ciphertext) == 0 {
		return nil, fmt.Errorf("the request must include the ciphertext to decrypt")
	}

	keyMetadata, err := e.Storage.GetKey(req.KeyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State == Disabled && !e.Cfg.AllowDecryptDisabled {
		return nil, fmt.Errorf("the key is disabled and decryption is not allowed for this key")
	}

	keyVersion, err := e.Storage.GetVersion(keyMetadata.KeyID, req.Version)
	if err != nil {
		return nil, err
	}

	plaintext, err := e.Crypto.Decrypt(keyMetadata.Algorithm, keyVersion.Material, req.Ciphertext, req.AdditionalData)
	if err != nil {
		return nil, err
	}

	response := DecryptResponse{Plaintext: plaintext}

	return &response, nil
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
