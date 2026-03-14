package main

import (
	"fmt"

	"github.com/michaeljmartin28/minikms/internal/config"
)

func main() {

	cfg := config.Load()

	fmt.Printf("Loaded config: %+v\n", cfg)
}