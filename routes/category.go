package routes

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/controller"
)

func SetupCategoryRoutes(router fiber.Router, categoryController controller.CategoryController) {
	router.Get("/api/categories", categoryController.FindAllByCategory)
	router.Post("/api/categories", categoryController.Create)
	router.Get("/api/categories/info", categoryController.SummaryCategory)
}
