package main

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"koriebruh/management/cnf"
	"koriebruh/management/controller"
	"koriebruh/management/repository"
	"koriebruh/management/service"
	"log"
)

func main() {
	db := cnf.InitDB()
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(db, authRepository)
	authController := controller.NewAuthController(authService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{}))

	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)
	app.Post("/api/auth/logout", authController.Logout)

	authorized := app.Group("/", cnf.JWTAuthMiddleware)
	authorized.Get("/hi", hellobg)

	authorized.Get("api/categories", categoryController.FindAllByCategory)
	authorized.Post("api/categories", categoryController.Create)

	err := app.Listen(cnf.GetConfig().Server.Port)
	if err != nil {
		log.Fatal("server terminated")
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("hello bg")
}
