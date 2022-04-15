package main

import (
	"github.com/hexhoc/go-mall-api/config"
	"github.com/hexhoc/go-mall-api/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
