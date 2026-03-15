package core

type Algorithm string

const (
	AES256GCM Algorithm = "AES-256-GCM"
)

type Ciphertext struct {
	KeyID     string
	Version   int
	Algorithm Algorithm
	Data      []byte
}

type Crypto interface {
	GenerateKey(algorithm Algorithm) ([]byte, error)
	Encrypt(alg Algorithm, key []byte, plaintext []byte, additionalData []byte) ([]byte, error)
	Decrypt(alg Algorithm, key []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
}