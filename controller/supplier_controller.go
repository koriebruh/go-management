package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"net/http"
)

type SupplierController interface {
	Create(ctx *fiber.Ctx) error
	FindAllSupplier(ctx *fiber.Ctx) error
	SummarySupplier(ctx *fiber.Ctx) error
}

type SupplierControllerImpl struct {
	service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) *SupplierControllerImpl {
	return &SupplierControllerImpl{SupplierService: supplierService}
}

func (controller SupplierControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.SupplierRequest
	token := ctx.Cookies("token")

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	err := controller.SupplierService.Create(ctx.Context(), token, request)
	if err != nil { // <-- if got error in service or repo
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: map[string]string{
			"message": "success created new supplier",
		},
	})
}

func (controller SupplierControllerImpl) FindAllSupplier(ctx *fiber.Ctx) error {
	suppliers, err := controller.SupplierService.FindAllSupplier(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   suppliers,
	})

}

func (controller SupplierControllerImpl) SummarySupplier(ctx *fiber.Ctx) error {
	summary, err := controller.SupplierService.SupplierSummary(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   summary,
	})

}
