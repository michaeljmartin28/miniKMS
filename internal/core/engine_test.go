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

func TestGenerateAndDecryptDataKey(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	eng, cleanup := newTestEngine(t)
	defer cleanup()

	// Create KEK
	createResp, err := eng.CreateKey(ctx, core.CreateKeyRequest{
		Name:      "kek",
		Algorithm: core.AES256GCM,
	})
	if err != nil {
		t.Fatalf("CreateKey: %v", err)
	}

	// Generate DEK
	aad := []byte("test-aad")
	genResp, err := eng.GenerateDataKey(ctx, core.GenerateDataKeyRequest{
		KeyID:          createResp.KeyID,
		AdditionalData: aad,
	})
	if err != nil {
		t.Fatalf("GenerateDataKey: %v", err)
	}

	if len(genResp.PlaintextDEK) != 32 {
		t.Fatalf("expected 32-byte DEK, got %d", len(genResp.PlaintextDEK))
	}
	if len(genResp.EncryptedDEK) == 0 {
		t.Fatalf("expected non-empty EncryptedDEK")
	}

	// Decrypt DEK
	decResp, err := eng.DecryptDataKey(ctx, core.DecryptDataKeyRequest{
		KeyID:          createResp.KeyID,
		EncryptedDEK:   genResp.EncryptedDEK,
		AdditionalData: aad,
		Version:        genResp.Version,
	})
	if err != nil {
		t.Fatalf("DecryptDataKey: %v", err)
	}

	// Assert equality
	if !bytes.Equal(decResp.PlaintextDEK, genResp.PlaintextDEK) {
		t.Fatalf("DEK mismatch")
	}
}

func TestDecryptDataKey_WrongAADFails(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	eng, cleanup := newTestEngine(t)
	defer cleanup()

	createResp, _ := eng.CreateKey(ctx, core.CreateKeyRequest{
		Name:      "kek",
		Algorithm: core.AES256GCM,
	})

	genResp, _ := eng.GenerateDataKey(ctx, core.GenerateDataKeyRequest{
		KeyID:          createResp.KeyID,
		AdditionalData: []byte("good"),
	})

	_, err := eng.DecryptDataKey(ctx, core.DecryptDataKeyRequest{
		KeyID:          createResp.KeyID,
		EncryptedDEK:   genResp.EncryptedDEK,
		AdditionalData: []byte("bad"),
		Version:        genResp.Version,
	})

	if err == nil {
		t.Fatalf("expected error when decrypting with wrong AAD")
	}
}

// func TestDecryptDataKey_DisabledKeyBehavior(t *testing.T) {
//     t.Parallel()

//     ctx := context.Background()
//     eng, cleanup := newTestEngine(t)
//     defer cleanup()

//     createResp, _ := eng.CreateKey(ctx, core.CreateKeyRequest{
//         Name:      "kek",
//         Algorithm: core.AES256GCM,
//     })

//     genResp, _ := eng.GenerateDataKey(ctx, core.GenerateDataKeyRequest{
//         KeyID: createResp.KeyID,
//     })

//     // Disable key
//     eng.Storage.UpdateKeyState(createResp.KeyID, core.KeyStateDisabled)

//     t.Run("disallowed", func(t *testing.T) {
//         eng.Cfg.AllowDecryptDisabled = false

//         _, err := eng.DecryptDataKey(ctx, core.DecryptDataKeyRequest{
//             KeyID:        createResp.KeyID,
//             EncryptedDEK: genResp.EncryptedDEK,
//             Version:      genResp.Version,
//         })

//         if err == nil {
//             t.Fatalf("expected error when decrypting with disabled key")
//         }
//     })

//     t.Run("allowed", func(t *testing.T) {
//         eng.Cfg.AllowDecryptDisabled = true

//         _, err := eng.DecryptDataKey(ctx, core.DecryptDataKeyRequest{
//             KeyID:        createResp.KeyID,
//             EncryptedDEK: genResp.EncryptedDEK,
//             Version:      genResp.Version,
//         })

//         if err != nil {
//             t.Fatalf("unexpected error: %v", err)
//         }
//     })
// }
