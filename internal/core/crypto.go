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
	Encrypt(alg Algorithm, plaintext []byte, key []byte) ([]byte, error)
	Decrypt(alg Algorithm, ciphertext []byte, key []byte) ([]byte, error)
}