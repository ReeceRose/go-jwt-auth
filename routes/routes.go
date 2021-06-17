package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
}
