package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/controllers"
)

// Setup routes which are available regardless of whether authenticated or not.
func SetupPublicRoutes(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
}

// Setup routes which require JWT authentication
func SetupPrivateRoutes(app *fiber.App) {

}
