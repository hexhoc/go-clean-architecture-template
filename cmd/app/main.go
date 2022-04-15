package main

import (
	"fmt"
	"github.com/hexhoc/go-mall-api/config"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Println(cfg.HTTP.Version, cfg.App.Version)
	// Run
	//app.Run(cfg)
}
