package crypto

type AESGCMCrypto struct{}

func NewAESGCMCrypto() *AESGCMCrypto {
	return &AESGCMCrypto{}
}

func (a *AESGCMCrypto) GenerateKey() ([]byte, error) {
	// implementation to generate a random 256-bit key
	return nil, nil
}

func (a *AESGCMCrypto) Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	// implementation for AES-GCM encryption
	return nil, nil
}

func (a *AESGCMCrypto) Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	// implementation for AES-GCM decryption
	return nil, nil
}
