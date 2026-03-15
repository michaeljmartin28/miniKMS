package storage

import (
	"os"
	"testing"
)

func NewTestStore(t *testing.T) (*BoltStore, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "testdb-*.bolt")
	if err != nil {
		t.Fatal(err)
	}

	store, err := NewBoltStore(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		store.Close()
		os.Remove(tmpFile.Name())
	}

	return store, cleanup
}
