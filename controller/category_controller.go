package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"koriebruh/management/utils"
	"net/http"
)

type CategoryController interface {
	Create(ctx *fiber.Ctx) error
	FindAllByCategory(ctx *fiber.Ctx) error
	SummaryCategory(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (controller CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var request dto.CategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	err := controller.CategoryService.Create(ctx, request)
	if err != nil { // <-- if got error in service or repo
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	// Respons sukses
	res := utils.SuccessRes(http.StatusCreated, "CREATED", "CREATED NEW CATEGORY SUCCESS")
	return ctx.Status(http.StatusCreated).JSON(res)

}

func (controller CategoryControllerImpl) FindAllByCategory(ctx *fiber.Ctx) error {
	category, err := controller.FindAllCategory(ctx)
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	// Response success
	res := utils.SuccessRes(http.StatusOK, "OK", category)
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller CategoryControllerImpl) SummaryCategory(ctx *fiber.Ctx) error {
	summary, err := controller.CategoryService.SummaryCategory(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", summary)
	return ctx.Status(http.StatusOK).JSON(res)
}
