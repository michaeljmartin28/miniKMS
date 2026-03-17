package core_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/michaeljmartin28/minikms/internal/core"
	"github.com/michaeljmartin28/minikms/internal/crypto"
	"github.com/michaeljmartin28/minikms/internal/storage"
)

func newTestEngine(t *testing.T) (*core.Engine, func()) {
	store, cleanup := storage.NewTestStore(t)
	crypto := crypto.NewAESGCMCrypto()
	cfg := core.DefaultConfig()

	engine := core.NewEngine(store, crypto, cfg)
	return engine, cleanup
}

func TestCreateKey(t *testing.T) {
	engine, cleanup := newTestEngine(t)
	defer cleanup()

	keyResponse, err := engine.CreateKey(context.Background(), core.CreateKeyRequest{Name: "test-key", Algorithm: core.AES256GCM})
	if err != nil {
		t.Fatalf("CreateKey failed: %v", err)
	}

	meta, err := engine.Storage.GetKey(keyResponse.KeyID)
	if err != nil {
		t.Fatalf("GetKey failed: %v", err)
	}

	if meta.LatestVersion != 1 {
		t.Fatalf("expected LatestVersion=1, got %d", meta.LatestVersion)
	}

	if meta.Algorithm != core.AES256GCM {
		t.Fatalf("expected Algorithm=%s, got %s", core.AES256GCM, meta.Algorithm)
	}

	if meta.State != core.KeyStateEnabled {
		t.Fatalf("expected State=%s, got %s", core.KeyStateEnabled, meta.State)
	}

}

func TestEncryptDecrypt(t *testing.T) {
	engine, cleanup := newTestEngine(t)
	defer cleanup()

	keyResponse, err := engine.CreateKey(context.Background(),
		core.CreateKeyRequest{
			Name:      "test-key",
			Algorithm: core.AES256GCM,
		},
	)
	if err != nil {
		t.Fatalf("createKey failed: %v", err)
	}
	if keyResponse.KeyID == "" {
		t.Fatalf("expected non-empty keyID")
	}

	plaintext := []byte("testDataToEncrypt")

	// Encrypt the plaintext
	encryptResponse, err := engine.Encrypt(
		context.Background(),
		core.EncryptRequest{
			KeyID:          keyResponse.KeyID,
			Plaintext:      plaintext,
			AdditionalData: nil,
		},
	)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}

	if bytes.Equal(encryptResponse.Ciphertext, plaintext) {
		t.Fatalf("ciphertext should not equal plaintext")
	}
	if encryptResponse.Version != 1 {
		t.Fatalf("expected version 1, got %d", encryptResponse.Version)
	}

	// Decrypt the ciphertext
	decryptResponse, err := engine.Decrypt(
		context.Background(),
		core.DecryptRequest{
			KeyID:          keyResponse.KeyID,
			Ciphertext:     encryptResponse.Ciphertext,
			AdditionalData: nil, Version: encryptResponse.Version,
		},
	)
	if err != nil {
		t.Fatalf("decrypt failed: %v", err)
	}

	// Ensure they match
	if !bytes.Equal(decryptResponse.Plaintext, plaintext) {
		t.Fatalf("expected decrypted=%s, got %s", plaintext, decryptResponse.Plaintext)
	}
}
