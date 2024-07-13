package main

import (
	"log"

	"order-service/api"
	"order-service/internal/config"
	"order-service/internal/service"
	"order-service/internal/storage"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs, nil)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*storage))

	log.Fatal(api.RUN(configs))
}
