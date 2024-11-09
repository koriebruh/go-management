package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
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
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	err := controller.CategoryService.Create(ctx, request)
	if err != nil { // <-- if got error in service or repo
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	// Respons sukses
	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: map[string]string{
			"message": "success created new category",
		},
	})

}

func (controller CategoryControllerImpl) FindAllByCategory(ctx *fiber.Ctx) error {
	category, err := controller.FindAllCategory(ctx)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	// Response success
	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	})

}

func (controller CategoryControllerImpl) SummaryCategory(ctx *fiber.Ctx) error {
	summary, err := controller.CategoryService.SummaryCategory(ctx.Context())
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
