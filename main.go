package main

import (
	"go-kafka-project/configs"
	"go-kafka-project/internal/api"
	"log"
)

func main() {

	cfg, err := configs.SetupEnv()

	if err != nil {
		log.Fatalf("something went wrong while setting up env: %v", err)
	}

	api.StartServer(cfg)
}
