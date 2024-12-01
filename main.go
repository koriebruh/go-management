package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"koriebruh/management/cnf"
	"koriebruh/management/dependency_injector"
	"koriebruh/management/routes"
	"log"
	"time"
)

func main() {

	authController, _ := dependency_injector.InitializeAuth()
	categoryController, _ := dependency_injector.InitializeCategory()
	itemController, _ := dependency_injector.InitializeItem()
	supplierController, _ := dependency_injector.InitializeSupplier()

	app := fiber.New(fiber.Config{
		BodyLimit:    10 * 1024 * 1024, // 10MB
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
		MaxAge:           86400,
	}))

	// Unauthorized route
	app.Get("/hi", hellobg)
	routes.SetupAuthRoutes(app, authController)

	// Authorized routes
	authorized := app.Group("/", cnf.JWTAuthMiddleware)
	routes.SetupCategoryRoutes(authorized, categoryController)
	routes.SetupItemRoutes(authorized, itemController)
	routes.SetupSupplierRoutes(authorized, supplierController)

	err := app.Listen(cnf.GetConfig().Server.Port)
	if err != nil {
		log.Fatal("server terminated")
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("hello bg")
}
