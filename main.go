package main

import (
	"go-microservice-project/configs"
	"go-microservice-project/internal/api"
	"log"
)

func main() {

	cfg, err := configs.SetupEnv()

	if err != nil {
		log.Fatalf("something went wrong while setting up env: %v", err)
	}

	api.StartServer(cfg)
}
