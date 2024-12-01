package main

import "C"
import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"koriebruh/management/cnf"
	"koriebruh/management/controller"
	"koriebruh/management/repository"
	"koriebruh/management/service"
	"log"
	"time"
)

func main() {
	db := cnf.InitDB()
	validate := validator.New()
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(db, authRepository, validate)
	authController := controller.NewAuthController(authService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	itemRepository := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepository, db, validate)
	itemController := controller.NewItemController(itemService)

	supplierRepository := repository.NewSupplierRepository()
	supplierService := service.NewSupplierService(db, supplierRepository, validate)
	supplierController := controller.NewSupplierController(supplierService)

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
		ExposeHeaders:    "Set-Cookie", // Tambahkan ini
		MaxAge:           86400,        // Tambahkan ini
	}))

	app.Get("/hi", hellobg)
	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)
	app.Post("/api/auth/logout", authController.Logout)

	authorized := app.Group("/", cnf.JWTAuthMiddleware)
	authorized.Get("/api/admins", authController.FindAllAdmin)
	authorized.Get("/hi", hellobg)

	authorized.Get("/api/categories", categoryController.FindAllByCategory)
	authorized.Post("/api/categories", categoryController.Create)
	authorized.Get("/api/categories/info", categoryController.SummaryCategory)

	authorized.Get("/api/items", itemController.FindAllByItem)
	authorized.Post("/api/items", itemController.CreateItem)
	authorized.Get("/api/items/info", itemController.SummaryItem)
	authorized.Get("/api/items/condition", itemController.FindByCondition)
	authorized.Get("/api/items/metric", itemController.InventoryMetrics)
	authorized.Get("/api/items/category", itemController.ReportItemByCategory)

	authorized.Post("/api/suppliers", supplierController.Create)
	authorized.Get("/api/suppliers", supplierController.FindAllSupplier)
	authorized.Get("/api/suppliers/info", supplierController.SummarySupplier)

	err := app.Listen(cnf.GetConfig().Server.Port)
	if err != nil {
		log.Fatal("server terminated")
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("hello bg")
}
