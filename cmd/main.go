package main

import (
	"log"

	"github.com/michaeljmartin28/minikms/internal/config"
	"github.com/michaeljmartin28/minikms/internal/storage"
)

func main() {

	cfg := config.Load()

	db, err := storage.NewBoltStore(cfg.DBPath)
	if err != nil{
		log.Fatalf("Error initializing BoltStore: %v\n", err)
		return
	}
	log.Println("Successfully initialized BoltStore")
	defer db.Close()

	log.Printf("Loaded config: %+v\n", cfg)

	
}