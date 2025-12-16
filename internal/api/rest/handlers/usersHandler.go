package handlers

import (
	"go-microservice-project/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// user service
}

func SetupUserHandler(rh *rest.RestHandler) {

	app := rh.App

	// TODO: create an instance of user service and inject to users handler
	userHandler := UserHandler{}

	// Create users group with /users prefix
	users := app.Group("/users")

	// public routes
	users.Post("/register", userHandler.Register)
	users.Post("/login", userHandler.Login)

	// private routes
	users.Get("/verify", userHandler.GetVerificationCode)
	users.Post("/verify", userHandler.Verify)
	users.Post("/profile", userHandler.CreatProfile)
	users.Get("/profile", userHandler.GetProfile)
	users.Post("/cart", userHandler.AddToCart)
	users.Get("/cart", userHandler.GetCarts)

	users.Get("/order", userHandler.GetOrders)
	users.Get("/order/:id", userHandler.GetOrder)

	users.Get("/become-seller", userHandler.BecomeSeller)
}

func (uh *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User registered successfully",
	})
}

func (uh *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User logined successfully",
	})
}

func (uh *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get verification code successfully",
	})
}

func (uh *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Code verified successfully",
	})
}

func (uh *UserHandler) CreatProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User profile created successfully",
	})
}

func (uh *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User profile fetched successfully",
	})
}

func (uh *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User add to cart item successfully",
	})
}

func (uh *UserHandler) GetCarts(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User get cart items successfully",
	})
}

func (uh *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User get order by id successfully",
	})
}

func (uh *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User get orders successfully",
	})
}

func (uh *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User become seller successfully",
	})
}
