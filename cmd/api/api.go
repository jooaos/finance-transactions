package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jooaos/pismo/internal/utils"
)

func main() {
	app := fiber.New()

	api := InitDependenciesApi()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(utils.NewApiSuccessResponse("Ok"))
	})

	app.Post("/accounts", api.Controllers.AccountController.Create)
	app.Get("/accounts/:id", api.Controllers.AccountController.GetById)
	app.Post("/transactions", api.Controllers.TransactionController.Create)

	app.Listen(":8080")
}
