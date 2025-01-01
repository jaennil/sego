package main

import (
	"log"
	"sego/config"
	"sego/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config.NewConfig error: %s", err)
	}

	app.Run(cfg)
}
