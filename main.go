package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/cache"
	"github.com/rpstvs/serverfibergo/router"
	"github.com/rpstvs/serverfibergo/sql"
)

func main() {

	db := sql.CreateDBInstance()
	cachedQuote := cache.CacheQuote{}
	app := fiber.New()
	router.SetupRoutes(app, db, &cachedQuote)
	app.Listen(":8080")
}
