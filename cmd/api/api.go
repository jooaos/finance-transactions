package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Printf("Intializing server...")

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON("Ok")
	})

	app.Listen(":8080")
}
