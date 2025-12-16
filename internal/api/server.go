package api

import (
	"fmt"
	"go-microservice-project/configs"
	"go-microservice-project/internal/api/rest"
	"go-microservice-project/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {

	addr := fmt.Sprintf(":%v", config.Port)

	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(addr)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserHandler(rh)
}
