package main

import (
	"nextclan/validator-register/mobile-validator-register-service/config"
	"nextclan/validator-register/mobile-validator-register-service/internal/app"
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
