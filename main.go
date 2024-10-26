package main

import (
	"github.com/gofiber/fiber/v2"
)

var cachedQuote *Response

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":8080")
}
