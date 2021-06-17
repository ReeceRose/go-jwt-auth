package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/database"
	"github.com/reecerose/go-jwt-auth/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
