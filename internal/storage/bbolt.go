package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/michaeljmartin28/minikms/internal/core"
	"go.etcd.io/bbolt"
)

type BoltStore struct {
	db *bbolt.DB
}

const (
	keyMetaBucket     = "KeyMetadata"
	keyVersionsBucket = "KeyVersions"
)

func NewBoltStore(dbPath string) (*BoltStore, error) {

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	// Open or create the BoltDB database
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	// Initialize the bucket for storing key metadata if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte(keyMetaBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	// Initialize the bucket for storing key versions if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte(keyVersionsBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return &BoltStore{db: db}, nil
}

func (s *BoltStore) Close() error {
	return s.db.Close()
}

func serializeKeyMetadata(meta core.KeyMetadata) ([]byte, error) {
	return json.Marshal(meta)

}

func deserializeKeyMetadata(data []byte) (core.KeyMetadata, error) {
	var meta core.KeyMetadata
	err := json.Unmarshal(data, &meta)
	if err != nil {
		return core.KeyMetadata{}, err
	}
	return meta, nil
}

func serializeKeyVersion(version core.KeyVersion) ([]byte, error) {
	return json.Marshal(version)
}

func deserializeKeyVersion(data []byte) (core.KeyVersion, error) {
	var version core.KeyVersion
	err := json.Unmarshal(data, &version)
	if err != nil {
		return core.KeyVersion{}, err
	}
	return version, nil
}

// compile-time assertion to ensure BoltStore implements the core.keystore interface
var _ core.KeyStore = (*BoltStore)(nil)

func (s *BoltStore) SaveKey(keyMeta core.KeyMetadata) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyMetaBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", keyMetaBucket)
		}
		// Serialize the KeyMetadata and store it using the keyID as the key
		keyID := []byte(keyMeta.KeyID)
		data, err := serializeKeyMetadata(keyMeta)
		if err != nil {
			return err
		}

		// Check if key already exists
		if bucket.Get(keyID) != nil {
			return fmt.Errorf("Key with ID %s already exists", keyMeta.KeyID)
		}

		return bucket.Put(keyID, data)
	})
}

func (s *BoltStore) GetKey(keyID string) (core.KeyMetadata, error) {

	var result core.KeyMetadata

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyMetaBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", keyMetaBucket)
		}

		val := bucket.Get([]byte(keyID))
		if val == nil {
			return fmt.Errorf("key %s not found", keyID)
		}
		meta, err := deserializeKeyMetadata(val)
		if err != nil {
			return err
		}
		result = meta
		return nil
	})
	if err != nil {
		return core.KeyMetadata{}, err
	}
	return result, nil
}

func (s *BoltStore) UpdateKey(keyMeta core.KeyMetadata) error {

	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyMetaBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", keyMetaBucket)
		}
		keyID := []byte(keyMeta.KeyID)
		// Make sure the key does already exist before updating
		if bucket.Get(keyID) == nil {
			return fmt.Errorf("Key with ID %s does not exist", keyMeta.KeyID)
		}
		data, err := serializeKeyMetadata(keyMeta)
		if err != nil {
			return err
		}
		return bucket.Put(keyID, data)
	})
}

func (s *BoltStore) DeleteKey(keyID string) error {

	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyMetaBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", keyMetaBucket)
		}
		return bucket.Delete([]byte(keyID))
	})
}

func (s *BoltStore) SaveVersion(keyID string, version core.KeyVersion) error {

	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyVersionsBucket))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", keyVersionsBucket)
		}

		// Use a composite key of "keyID:version" to store each version
		compositeKey := fmt.Sprintf("%s:%d", keyID, version.Version)
		data, err := serializeKeyVersion(version)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(compositeKey), data)

	})

}

func (s *BoltStore) GetVersion(keyID string, version uint32) (core.KeyVersion, error) {

	var result core.KeyVersion

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyVersionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", keyVersionsBucket)
		}

		val := bucket.Get([]byte(fmt.Sprintf("%s:%d", keyID, version)))
		if val == nil {
			return fmt.Errorf("version %d for key %s not found", version, keyID)
		}
		meta, err := deserializeKeyVersion(val)
		if err != nil {
			return err
		}
		result = meta
		return nil
	})
	if err != nil {
		return core.KeyVersion{}, err
	}
	return result, nil
}

func (s *BoltStore) ListVersions(keyID string) ([]core.KeyVersion, error) {

	versions := make([]core.KeyVersion, 0, 4)

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(keyVersionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", keyVersionsBucket)
		}

		c := bucket.Cursor()
		prefix := []byte(fmt.Sprintf("%s:", keyID))
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			version, err := deserializeKeyVersion(v)
			if err != nil {
				return err
			}
			versions = append(versions, version)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return versions, nil
}
