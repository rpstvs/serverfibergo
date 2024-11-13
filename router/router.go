package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rpstvs/serverfibergo/cache"
	"github.com/rpstvs/serverfibergo/database"
	"github.com/rpstvs/serverfibergo/handler"
)

func SetupRoutes(app *fiber.App, db *database.Queries, quoteCache *cache.CacheQuote) {
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/author", func(c *fiber.Ctx) error {
		log.Println("Retreiving author info")
		return handler.GetAuthor(c, quoteCache)
	})
	quote := app.Group("/quote", logger.New())

	quote.Get("/random", func(c *fiber.Ctx) error {
		return handler.GetRandomQuote(c, db, quoteCache)
	})
}
