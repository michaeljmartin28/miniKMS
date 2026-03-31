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

func DefaultConfig() EngineConfig {
	return EngineConfig{
		DefaultAlgorithm:     AES256GCM,
		AllowDecryptDisabled: false,
		MaxKeyVersions:       0,
	}
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

func (e *Engine) CreateKey(ctx context.Context, req CreateKeyRequest) (*KeyMetadata, error) {

	keyBytes, err := e.Crypto.GenerateKey(req.Algorithm)
	version := uint32(1)

	keyVersion := KeyVersion{
		Version:   version,
		CreatedAt: time.Now(),
		Material:  keyBytes,
	}

	meta := KeyMetadata{
		KeyID:         uuid.New().String(),
		CreatedAt:     keyVersion.CreatedAt,
		Algorithm:     req.Algorithm,
		State:         KeyStateEnabled,
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

	return &meta, nil
}

func (e *Engine) Encrypt(ctx context.Context, req EncryptRequest) (*EncryptResponse, error) {

	keyMetadata, err := e.Storage.GetKey(req.KeyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State.IsDisabled() {
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

	keyMetadata, err := e.Storage.GetKey(req.KeyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State.IsDisabled() && !e.Cfg.AllowDecryptDisabled {
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

	keyMetadata, err := e.Storage.GetKey(req.KeyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State.IsDisabled() {
		return nil, fmt.Errorf("key is currently disabled and cannot be used to create a DEK")
	}

	// DEK uses the same algorithm as the KEK
	dek, err := e.Crypto.GenerateKey(keyMetadata.Algorithm)
	if err != nil {
		return nil, err
	}

	encryptResp, err := e.Encrypt(
		ctx,
		EncryptRequest{
			KeyID:     keyMetadata.KeyID,
			Plaintext: dek, AdditionalData: req.AdditionalData,
		},
	)
	if err != nil {
		return nil, err
	}

	response := GenerateDataKeyResponse{
		PlaintextDEK: dek,
		EncryptedDEK: encryptResp.Ciphertext,
		Version:      encryptResp.Version,
	}

	return &response, nil
}

func (e *Engine) DecryptDataKey(ctx context.Context, req DecryptDataKeyRequest) (*DecryptDataKeyResponse, error) {

	decResp, err := e.Decrypt(
		ctx,
		DecryptRequest{
			KeyID:          req.KeyID,
			Ciphertext:     req.EncryptedDEK,
			AdditionalData: req.AdditionalData,
			Version:        req.Version,
		},
	)
	if err != nil {
		return nil, err
	}

	response := DecryptDataKeyResponse{
		PlaintextDEK: decResp.Plaintext,
	}

	return &response, nil
}

func (e *Engine) RotateKey(ctx context.Context, keyID string) (uint32, error) {

	// Get the metadata for a key
	keyMetadata, err := e.Storage.GetKey(keyID)
	if err != nil {
		return 0, err
	}

	// Ensure we are allowed to rotate it
	if keyMetadata.State.IsDisabled() {
		return 0, fmt.Errorf("key is disabled - cannot rotate")
	}
	// TODO: maybe add another state, pending deletion

	// TODO: Add maxVersion controls. (0 == infinite, otherwise if we hit the max, delete the oldest or reject the request.)

	// Generate new key material
	newKey, err := e.Crypto.GenerateKey(keyMetadata.Algorithm)
	if err != nil {
		return 0, err
	}

	// Create a new version
	newVersion := KeyVersion{
		Version:   keyMetadata.LatestVersion + 1,
		CreatedAt: time.Now(),
		Material:  newKey,
	}

	// Store the new version
	err = e.Storage.SaveVersion(keyMetadata.KeyID, newVersion)
	if err != nil {
		return 0, err
	}

	// Now that we were able to store the new version, update the key metadata to match the newest version
	keyMetadata.LatestVersion++

	err = e.Storage.UpdateKey(keyMetadata)
	if err != nil {
		return 0, err
	}

	return keyMetadata.LatestVersion, nil

}

func (e *Engine) DisableKey(ctx context.Context, keyID string) (*KeyMetadata, error) {

	keyMetadata, err := e.Storage.GetKey(keyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State.IsDisabled() {
		return nil, fmt.Errorf("key is already disabled")
	}

	keyMetadata.State = KeyStateDisabled
	err = e.Storage.UpdateKey(keyMetadata)
	if err != nil {
		return nil, err
	}
	return &keyMetadata, nil
}

func (e *Engine) EnableKey(ctx context.Context, keyID string) (*KeyMetadata, error) {
	keyMetadata, err := e.Storage.GetKey(keyID)
	if err != nil {
		return nil, err
	}

	if keyMetadata.State.IsEnabled() {
		return nil, fmt.Errorf("key is already enabled")
	}
	keyMetadata.State = KeyStateEnabled
	err = e.Storage.UpdateKey(keyMetadata)
	if err != nil {
		return nil, err
	}
	return &keyMetadata, nil
}
