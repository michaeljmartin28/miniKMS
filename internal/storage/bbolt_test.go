package storage

import (
	"bytes"
	"os"
	"testing"

	"github.com/michaeljmartin28/minikms/internal/core"
)

func newTestStore(t *testing.T) *BoltStore {
    t.Helper()

    tmpFile, err := os.CreateTemp("", "testdb-*.bolt")
    if err != nil {
        t.Fatal(err)
    }

    store, err := NewBoltStore(tmpFile.Name())
    if err != nil {
        t.Fatal(err)
    }

    return store
}

func TestSaveAndGetKey(t *testing.T) {
    store := newTestStore(t)

    meta := core.KeyMetadata{
        KeyID:         "test-key",
        LatestVersion: 1,
    }

    if err := store.SaveKey(meta); err != nil {
        t.Fatalf("SaveKey failed: %v", err)
    }

    got, err := store.GetKey("test-key")
    if err != nil {
        t.Fatalf("GetKey failed: %v", err)
    }

    if got.KeyID != meta.KeyID {
        t.Errorf("expected ID %s, got %s", meta.KeyID, got.KeyID)
    }
}

func TestSaveAndGetVersion(t *testing.T) {
    store := newTestStore(t)

    v := core.KeyVersion{
        Version: 1,
        Material: []byte("abc123"),
    }

    if err := store.SaveVersion("key1", v); err != nil {
        t.Fatalf("SaveVersion failed: %v", err)
    }

    got, err := store.GetVersion("key1", 1)
    if err != nil {
        t.Fatalf("GetVersion failed: %v", err)
    }

    if !bytes.Equal(got.Material, v.Material) {
        t.Errorf("expected %v, got %v", v.Material, got.Material)
    }
}

func TestListVersions(t *testing.T) {
    store := newTestStore(t)

    store.SaveVersion("key1", core.KeyVersion{Version: 1})
    store.SaveVersion("key1", core.KeyVersion{Version: 2})

    versions, err := store.ListVersions("key1")
    if err != nil {
        t.Fatalf("ListVersions failed: %v", err)
    }

    if len(versions) != 2 {
        t.Fatalf("expected 2 versions, got %d", len(versions))
    }

    if versions[0].Version != 1 || versions[1].Version != 2 {
        t.Errorf("versions not in order: %+v", versions)
    }
}