package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	kmsv1 "github.com/michaeljmartin28/minikms/gen/kms/v1"
	"github.com/michaeljmartin28/minikms/internal/config"
	"github.com/michaeljmartin28/minikms/internal/core"
	"github.com/michaeljmartin28/minikms/internal/crypto"
	"github.com/michaeljmartin28/minikms/internal/storage"
	"github.com/michaeljmartin28/minikms/internal/transport/grpcsrv"
	"github.com/michaeljmartin28/minikms/internal/transport/httpsrv"
	"github.com/michaeljmartin28/minikms/package/version"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("miniKMS %s starting...\n", version.Version)
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

	// Set up the HTTP server

	httpHandler := httpsrv.NewHandler(engine)
	httpMux := httpsrv.NewRouter(httpHandler)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: httpMux,
	}

	// Set up the gRPC Server

	grpcServer := grpc.NewServer()
	kmsv1.RegisterKMSServer(grpcServer, grpcsrv.NewGRPCServer(engine))
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("grpc server error: %v", err)
		}
	}()

	func() {
		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		<-sigs
	}()

	httpServer.Shutdown(context.Background())
	grpcServer.GracefulStop()

}
