package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reecerose/go-jwt-auth/controllers"
)

// SetupPublicRoutes is used to register routes which require no authentication.
func SetupPublicRoutes(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
}

// SetupPrivateRoutes is used to register routes which require JWT authentication.
func SetupPrivateRoutes(app *fiber.App) {

}
