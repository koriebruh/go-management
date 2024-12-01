package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"koriebruh/management/utils"
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
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	err := controller.SupplierService.Create(ctx.Context(), token, request)
	if err != nil { // <-- if got error in service or repo
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusCreated, "SUCCESS", map[string]string{"message": "success created new supplier"})
	return ctx.Status(http.StatusCreated).JSON(res)
}

func (controller SupplierControllerImpl) FindAllSupplier(ctx *fiber.Ctx) error {
	suppliers, err := controller.SupplierService.FindAllSupplier(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", suppliers)
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller SupplierControllerImpl) SummarySupplier(ctx *fiber.Ctx) error {
	summary, err := controller.SupplierService.SupplierSummary(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", summary)
	return ctx.Status(http.StatusOK).JSON(res)

}
