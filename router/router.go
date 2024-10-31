package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rpstvs/serverfibergo/database"
	"github.com/rpstvs/serverfibergo/handler"
)

func SetupRoutes(app *fiber.App, db *database.Queries) {
	app.Use(cors.New())

	quote := app.Group("/quote", logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	quote.Get("/random", func(c *fiber.Ctx) error {
		return handler.GetRandomQuote(c, db)
	})
}
