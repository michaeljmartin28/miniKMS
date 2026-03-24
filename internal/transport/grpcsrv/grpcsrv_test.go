package grpcsrv_test

import (
	"context"
	"testing"

	kmsv1 "github.com/michaeljmartin28/minikms/gen/kms/v1"
	"github.com/michaeljmartin28/minikms/internal/core"
	"github.com/michaeljmartin28/minikms/internal/crypto"
	"github.com/michaeljmartin28/minikms/internal/storage"
	"github.com/michaeljmartin28/minikms/internal/transport/grpcsrv"
)

func TestEncrypt(t *testing.T) {
	storage, cleanup := storage.NewTestStore(t)
	cfg := core.DefaultConfig()
	crypto := crypto.NewAESGCMCrypto()
	engine := core.NewEngine(storage, crypto, cfg)
	server := grpcsrv.NewGRPCServer(engine)

	// Arrange: create a key
	createResp, _ := server.CreateKey(context.Background(), &kmsv1.CreateKeyRequest{
		Algorithm: string(core.AES256GCM),
	})

	// Act: encrypt something
	resp, err := server.Encrypt(context.Background(), &kmsv1.EncryptRequest{
		KeyId:     createResp.KeyId,
		Plaintext: []byte("hello"),
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Ciphertext) == 0 {
		t.Fatalf("expected ciphertext")
	}

	cleanup()
}
