package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/reecerose/go-jwt-auth/database"
	"github.com/reecerose/go-jwt-auth/routes"
	"github.com/reecerose/go-jwt-auth/utils"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.SECRET_KEY),
	}))

	app.Listen(":3000")
}
