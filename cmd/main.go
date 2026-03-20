package main

import (
	"log"
	netHttp "net/http"

	"github.com/michaeljmartin28/minikms/internal/config"
	"github.com/michaeljmartin28/minikms/internal/core"
	"github.com/michaeljmartin28/minikms/internal/crypto"
	"github.com/michaeljmartin28/minikms/internal/storage"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewBoltStore(cfg.DBPath)
	if err != nil {
		log.Fatalf("Error initializing BoltStore: %v\n", err)
		return
	}
	log.Println("Successfully initialized BoltStore")
	defer store.Close()

	log.Printf("Loaded config: %+v\n", cfg)

	crypto := crypto.NewAESGCMCrypto()

	engine := core.NewEngine(
		store,
		crypto,
		core.EngineConfig{
			DefaultAlgorithm:     core.AES256GCM,
			MaxKeyVersions:       10,
			AllowDecryptDisabled: false,
		},
	)

	log.Printf("Engine initialized with config: %+v\n", engine.Cfg)

	netHttp.ListenAndServe(":8080", nil)

}
