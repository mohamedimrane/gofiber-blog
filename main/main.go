package main

import "github.com/gofiber/fiber/v2"

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello")

		return nil
	})

	app.Listen(":8080")
}
