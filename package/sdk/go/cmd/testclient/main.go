package main

import (
	"context"
	"fmt"
	"log"

	kms "github.com/michaeljmartin28/minikms/package/sdk/go"
)

func main() {
	ctx := context.Background()
	client := kms.NewClient("http://localhost:8080")

	// 1. Create a key
	key, err := client.CreateKey(ctx, kms.CreateKeyParams{
		Algorithm: "AES-256-GCM",
		Name:      "sdk-test",
	})
	fmt.Printf("key: %+v\n", key)

	if err != nil {
		log.Fatalf("CreateKey failed: %v", err)
	}

	fmt.Println("Created key:", key.KeyID)

	// 2. Encrypt
	enc, err := client.Encrypt(ctx, key.KeyID, kms.EncryptParams{
		Plaintext: []byte("hello sdk"),
	})
	if err != nil {
		log.Fatalf("Encrypt failed: %v", err)
	}
	fmt.Println("Ciphertext:", enc.Ciphertext)

	// 3. Decrypt
	dec, err := client.Decrypt(ctx, key.KeyID, kms.DecryptParams{
		Ciphertext: enc.Ciphertext,
		Version:    key.Version,
	})
	if err != nil {
		log.Fatalf("Decrypt failed: %v", err)
	}
	fmt.Println("Decrypted:", string(dec.Plaintext))

	// 4. Generate a data key
	dk, err := client.GenerateDEK(ctx, key.KeyID, kms.GenerateDataParams{})
	if err != nil {
		log.Fatalf("GenerateDataKey failed: %v", err)
	}
	fmt.Println("Generated DEK (plaintext):", dk.PlaintextDEK)
	fmt.Println("Generated DEK (encrypted):", dk.EncryptedDEK)

	// 5. Decrypt the data key
	pt, err := client.DecryptDEK(ctx, key.KeyID, kms.DecryptDataKeyParams{
		EncryptedDEK: dk.EncryptedDEK,
		Version:      dk.Version,
	})
	if err != nil {
		log.Fatalf("DecryptDataKey failed: %v", err)
	}
	fmt.Println("Unwrapped DEK:", pt.PlaintextDEK)
}
