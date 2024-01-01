package handler

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"be-park-ease/middleware"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandler interface {
	AllHistory(ctx *fiber.Ctx) error
	CreateEntryHistory(ctx *fiber.Ctx) error
	CalculatePriceHistory(ctx *fiber.Ctx) error
}

type historyHandler struct {
	conf      *config.Config
	service   service.HistoryService
	exception exception.Exception
}

func NewHistoryHandler(service service.HistoryService) HistoryHandler {
	return &historyHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("history-handler"),
	}
}

// AllHistory godoc
//
//	@Summary		Get All History based on parameter
//	@Description	All History
//	@ID				get-all-history
//	@Tags			history
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page			query		int		false	"Page"	default(1)
//	@Param			limit			query		int		false	"Limit"	default(10)
//	@Param			search			query		string	false	"Search"
//	@Param			order_by		query		string	false	"Order by"	Enums(id,location_code,vehicle_type_code,vehicle_number,date,type)
//	@Param			order			query		string	false	"Order"		Enums(asc, desc)
//	@Param			type			query		string	false	"Type"		Enums(entry,exit,fine)
//	@Param			vehicle_type	query		string	false	"Vehicle Type"
//	@Param			location		query		string	false	"Location"
//	@Success		200				{object}	response.BaseResponse{data=response.BaseResponsePagination[response.EntryHistory]}
//	@Failure		500				{object}	response.BaseResponse
//	@Router			/history [get]
func (h *historyHandler) AllHistory(ctx *fiber.Ctx) error {
	req := request.AllHistoryRequest{
		BasePagination: request.BasePagination{
			Page:    ctx.QueryInt("page", 1),
			Limit:   ctx.QueryInt("limit", constants.DefaultPageLimit),
			Search:  ctx.Query("search"),
			OrderBy: ctx.Query("order_by"),
			Order:   ctx.Query("order"),
		},
		HistoryType: constants.HistoryType(ctx.Query("type")),
		VehicleType: ctx.Query("vehicle_type"),
		Location:    ctx.Query("location"),
	}

	fieldOrder := map[string]string{
		"id":                "eh.id",
		"location_code":     "eh.location_code",
		"vehicle_type_code": "eh.vehicle_type_code",
		"vehicle_number":    "eh.vehicle_number",
		"date":              "date",
		"type":              "type",
	}

	req.ValidateAndUpdateOrderBy(fieldOrder)
	req.Normalize()

	res := h.service.AllHistory(ctx.Context(), req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// CreateEntryHistory godoc
//
//	@Summary		Post Create Entry History based on parameter
//	@Description	Create Entry History
//	@ID				post-create-entry-history
//	@Tags			history
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.CreateEntryHistoryRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/history/entry [post]
func (h *historyHandler) CreateEntryHistory(ctx *fiber.Ctx) error {
	req := new(request.CreateEntryHistoryRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID
	req.VehicleNumber = strings.ToUpper(req.VehicleNumber)

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.CreateEntryHistory(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// CalculatePriceHistory godoc
//
//	@Summary		Post Calculate Price History based on parameter
//	@Description	Calculate Price History
//	@ID				post-calculate-price-history
//	@Tags			history
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.CalculatePriceHistoryRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse{data=float64}
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/history/calculate [post]
func (h *historyHandler) CalculatePriceHistory(ctx *fiber.Ctx) error {
	req := new(request.CalculatePriceHistoryRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)
	req.VehicleNumber = strings.ToUpper(req.VehicleNumber)

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	res := h.service.CalculatePriceHistory(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}
