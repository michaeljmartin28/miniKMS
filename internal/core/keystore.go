package core

import "time"

// KeyMetadata represents the metadata associated with a key, including its unique identifier,
// creation time, and current state.
type KeyMetadata struct {
	KeyID         string    `json:"keyId"`
	CreatedAt     time.Time `json:"createdAt"`
	State         KeyState  `json:"state"`
	Algorithm     Algorithm `json:"algorithm"`
	LatestVersion uint32    `json:"latestVersion"`
}

// KeyVersion represents a specific version of a key, including its version number,
// creation time, and the key material.
type KeyVersion struct {
	Version   uint32    `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
	Material  []byte    `json:"material"`
}

//
type KeyState string

func (s KeyState) IsEnabled() bool {
	return s == KeyStateEnabled
}

func (s KeyState) IsDisabled() bool {
	return s == KeyStateDisabled
}

const (
	KeyStateEnabled  KeyState = "ENABLED"
	KeyStateDisabled KeyState = "DISABLED"
)

type KeyStore interface {
	// Metadata operations
	SaveKey(meta KeyMetadata) error
	GetKey(id string) (KeyMetadata, error)
	UpdateKey(meta KeyMetadata) error
	DeleteKey(id string) error

	// Version operations
	SaveVersion(keyID string, version KeyVersion) error
	GetVersion(keyID string, version uint32) (KeyVersion, error)
	ListVersions(keyID string) ([]KeyVersion, error)
}
