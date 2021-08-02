package main

import (
	"log"

	"github.com/mozg1984/delivery_routing_service/pkg/api"
	"github.com/mozg1984/delivery_routing_service/pkg/app"
	"github.com/mozg1984/delivery_routing_service/pkg/app/repository"
	"github.com/mozg1984/delivery_routing_service/pkg/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Startup error '%s' was occurred", err)
	}
}

func run() error {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("An error '%s' was occurred while reading the configuration file", err)
	}

	storage := repository.NewStorage(&config.Redis)
	deliveryService := api.NewDeliveryService(storage)

	handler := app.NewHandler()
	router := app.NewRouter(handler)

	server := app.NewServer(router, deliveryService)
	server.Run(&config.Server)

	return nil
}
