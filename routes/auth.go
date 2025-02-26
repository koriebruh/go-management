package routes

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/cnf"
	"koriebruh/management/controller"
)

func SetupAuthRoutes(app *fiber.App, authController controller.AuthController) {
	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)

	authorized := app.Group("/", cnf.JWTAuthMiddleware)
	authorized.Post("/api/auth/logout", authController.Logout)
	authorized.Get("/api/admins", authController.FindAllAdmin)
}
