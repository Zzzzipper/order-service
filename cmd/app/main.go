package main

import (
	"log"

	"gitlab.mapcard.pro/external-map-team/order-service/config"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
