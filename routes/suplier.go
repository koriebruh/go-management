package routes

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/controller"
)

func SetupSupplierRoutes(router fiber.Router, supplierController controller.SupplierController) {
	router.Post("/api/suppliers", supplierController.Create)
	router.Get("/api/suppliers", supplierController.FindAllSupplier)
	router.Get("/api/suppliers/info", supplierController.SummarySupplier)
}
