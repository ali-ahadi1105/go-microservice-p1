package api

import (
	"fmt"
	"go-microservice-project/configs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {

	addr := fmt.Sprintf(":%v", config.Port)

	app := fiber.New()

	app.Get("/health", Healthcheck)

	app.Listen(addr)
}

func Healthcheck(ctx *fiber.Ctx) error {
	log.Println("Calling Healthcheck request")
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm breathing Bastards!",
	})
}
