package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/michaeljmartin28/minikms/internal/core"
)

type AESGCMCrypto struct{}

func NewAESGCMCrypto() *AESGCMCrypto {
	return &AESGCMCrypto{}
}

func (a *AESGCMCrypto) GenerateKey(alg core.Algorithm) ([]byte, error) {
	switch alg {
	case core.AES256GCM:
		key := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, key); err != nil {
			return nil, err
		}
		return key, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %v", alg)
	}
}

func (a *AESGCMCrypto) Encrypt(alg core.Algorithm, key []byte, plaintext []byte, add []byte) ([]byte, error) {
	switch alg {

	case core.AES256GCM:
		aes, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		// Create a GCM cipher mode instance
		gcm, err := cipher.NewGCM(aes)
		if err != nil {
			return nil, err
		}

		// Generate a random nonce
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil{
			return nil, err
		}

		ciphertext := gcm.Seal(nil, nonce, plaintext, add)
		ciphertext = append(nonce, ciphertext...)

		return ciphertext, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %v", alg)
	}
	
}

func (a *AESGCMCrypto) Decrypt(alg core.Algorithm, key []byte, ciphertext []byte, add []byte) ([]byte, error) {
	switch alg {
	case core.AES256GCM:
		// Create a new AES cipher block
		aes, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		// Create a GCM cipher mode instance
		gcm, err := cipher.NewGCM(aes)
		if err != nil {
			return nil, err
		}

		// Extract the nonce from the beginning of the ciphertext
		nonceSize := gcm.NonceSize()
		if len(ciphertext) < nonceSize {
			return nil, fmt.Errorf("ciphertext too short")
		}

		nonce := ciphertext[:nonceSize]

		// Decrypt the ciphertext using the nonce and additional data
		ciphertext = ciphertext[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, add)
		if err != nil {
			return nil, err
		}

		return plaintext, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %v", alg)
	}
}
