package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"koriebruh/management/utils"
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
	ReportItemByCategory(ctx *fiber.Ctx) error
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
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	err := controller.ItemService.Create(ctx.Context(), token, request)
	if err != nil { // <-- if got error in service or repo
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusCreated, "SUCCESS", map[string]string{"message": "success created new item"})
	return ctx.Status(http.StatusCreated).JSON(res)

}

func (controller ItemControllerImpl) FindAllByItem(ctx *fiber.Ctx) error {

	item, err := controller.ItemService.FindAllItem(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", item)
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller ItemControllerImpl) SummaryItem(ctx *fiber.Ctx) error {
	summary, err := controller.ItemService.SummaryItem(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", summary)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (controller ItemControllerImpl) FindByCondition(ctx *fiber.Ctx) error {
	// TAKE VALUE PARAM
	condition := ctx.Query("condition")
	thresholdStr := ctx.Query("threshold", "0")

	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		err = errors.New("error: Invalid threshold value")
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	if condition == "" {
		log.Println("error condition ")
		err := errors.New("error: Condition parameter is required")
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	valueCondition, err := controller.ItemService.FindByCondition(ctx.Context(), condition, threshold)
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", valueCondition)
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller ItemControllerImpl) InventoryMetrics(ctx *fiber.Ctx) error {
	metric, err := controller.ItemService.InventoryMetrics(ctx.Context())
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", metric)
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller ItemControllerImpl) ReportItemByCategory(ctx *fiber.Ctx) error {

	condition := ctx.Query("category")

	report, err := controller.ItemService.ReportItemByCategory(ctx.Context(), condition)
	if err != nil {
		responseWeb := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(responseWeb)
	}

	res := utils.SuccessRes(http.StatusOK, "OK", report)
	return ctx.Status(http.StatusOK).JSON(res)
}
