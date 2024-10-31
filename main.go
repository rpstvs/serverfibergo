package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/router"
	"github.com/rpstvs/serverfibergo/sql"
)

func main() {

	db := sql.CreateDBInstance()
	app := fiber.New()
	router.SetupRoutes(app, db)
	app.Listen(":8080")
}
