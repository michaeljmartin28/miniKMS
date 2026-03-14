package core

// KeyMetadata represents the metadata of a key, including its ID, versions, and creation time.
type KeyMetadata struct {
	KeyID     string
	Versions  []KeyVersion
	CreatedAt int64
}

// KeyVersion represents a specific version of a key, including its version number, 
// creation time, key bytes, and whether it is disabled.
type KeyVersion struct {
	Version   int
	CreatedAt int64
	KeyBytes []byte
	Disabled  bool
}

// KeyStore defines the interface for a key storage system, allowing for saving, 
// retrieving, and updating key metadata.
type KeyStore interface {
	SaveKey(keyMeta KeyMetadata) error
	GetKey(keyID string) (KeyMetadata, error)
	UpdateKey(keyMeta KeyMetadata) error
}