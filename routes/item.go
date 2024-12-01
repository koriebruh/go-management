package routes

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/controller"
)

func SetupItemRoutes(router fiber.Router, itemController controller.ItemController) {
	router.Get("/api/items", itemController.FindAllByItem)
	router.Post("/api/items", itemController.CreateItem)
	router.Get("/api/items/info", itemController.SummaryItem)
	router.Get("/api/items/condition", itemController.FindByCondition)
	router.Get("/api/items/metric", itemController.InventoryMetrics)
	router.Get("/api/items/category", itemController.ReportItemByCategory)
}
