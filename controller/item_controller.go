package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"log"
	"net/http"
	"strconv"
)

type ItemController interface {
	CreateItem(ctx *fiber.Ctx) error
	FindAllByItem(ctx *fiber.Ctx) error
	SummaryItem(ctx *fiber.Ctx) error
	FindByCondition(ctx *fiber.Ctx) error
	InventoryMetrics(ctx *fiber.Ctx) error
}

type ItemControllerImpl struct {
	service.ItemService
}

func NewItemController(itemService service.ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{ItemService: itemService}
}

func (controller ItemControllerImpl) CreateItem(ctx *fiber.Ctx) error {

	var request dto.ItemRequest
	token := ctx.Cookies("token")

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	err := controller.ItemService.Create(ctx.Context(), token, request)
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
			"message": "success created new item",
		},
	})

}

func (controller ItemControllerImpl) FindAllByItem(ctx *fiber.Ctx) error {

	item, err := controller.ItemService.FindAllItem(ctx.Context())
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
		Data:   item,
	})

}

func (controller ItemControllerImpl) SummaryItem(ctx *fiber.Ctx) error {
	summary, err := controller.ItemService.SummaryItem(ctx.Context())
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

func (controller ItemControllerImpl) FindByCondition(ctx *fiber.Ctx) error {
	// TAKE VALUE PARAM
	condition := ctx.Query("condition")
	thresholdStr := ctx.Query("threshold", "0")

	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		log.Printf("error thresholdStr %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: map[string]string{
				"error": "Invalid threshold value",
			},
		})
	}

	if condition == "" {
		log.Println("error condition ")
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: map[string]string{
				"error": "Condition parameter is required",
			},
		})
	}

	valueCondition, err := controller.ItemService.FindByCondition(ctx.Context(), condition, threshold)
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
		Data:   valueCondition,
	})

}

func (controller ItemControllerImpl) InventoryMetrics(ctx *fiber.Ctx) error {
	metric, err := controller.ItemService.InventoryMetrics(ctx.Context())
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
		Data:   metric,
	})
}
