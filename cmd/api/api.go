package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jooaos/pismo/internal/utils"
)

func main() {
	app := fiber.New()

	_ = InitDependenciesApi()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(utils.NewApiSuccessResponse("Ok"))
	})

	app.Listen(":8080")
}
