package api

import (
	"fmt"
	"go-kafka-project/configs"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {

	addr := fmt.Sprintf(":%v", config.Port)

	app := fiber.New()

	app.Listen(addr)
}
