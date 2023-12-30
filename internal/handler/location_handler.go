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

	"github.com/gofiber/fiber/v2"
)

type LocationHandler interface {
	AllLocation(ctx *fiber.Ctx) error
	LocationByCode(ctx *fiber.Ctx) error
	CreateLocation(ctx *fiber.Ctx) error
	UpdateLocation(ctx *fiber.Ctx) error
	DeleteLocation(ctx *fiber.Ctx) error
}

type locationHandler struct {
	conf      *config.Config
	service   service.LocationService
	exception exception.Exception
}

func NewLocationHandler(service service.LocationService) LocationHandler {
	return &locationHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("location-handler"),
	}
}

// AllLocation godoc
//
//	@Summary		Get All Location based on parameter
//	@Description	All Location
//	@ID				get-all-location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page		query		int		false	"Page"	default(1)
//	@Param			limit		query		int		false	"Limit"	default(10)
//	@Param			search		query		string	false	"Search"
//	@Param			order_by	query		string	false	"Order by"	Enums(code,name,is_exit,date)
//	@Param			order		query		string	false	"Order"		Enums(asc, desc)
//	@Success		200			{object}	response.BaseResponse{data=response.BaseResponsePagination[response.Location]}
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/location [get]
func (h *locationHandler) AllLocation(ctx *fiber.Ctx) error {
	req := request.BasePagination{
		Page:    ctx.QueryInt("page", 1),
		Limit:   ctx.QueryInt("limit", constants.DefaultPageLimit),
		Search:  ctx.Query("search"),
		OrderBy: ctx.Query("order_by"),
		Order:   ctx.Query("order"),
	}

	fieldOrder := map[string]string{
		"code":    "code",
		"name":    "name",
		"is_exit": "is_exit",
		"date":    "created_at",
	}

	req.ValidateAndUpdateOrderBy(fieldOrder)
	req.Normalize()

	res := h.service.AllLocation(ctx.Context(), req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// LocationByCode godoc
//
//	@Summary		Get Location based on code
//	@Description	Location
//	@ID				get-location-by-code
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			code	path		string	true	"Code"
//	@Success		200		{object}	response.BaseResponse{data=response.Location}
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/location/{code} [get]
func (h *locationHandler) LocationByCode(ctx *fiber.Ctx) error {
	res := h.service.LocationByCode(ctx.Context(), ctx.Params("code"))
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// CreateLocation godoc
//
//	@Summary		Create Location
//	@Description	Location
//	@ID				create-location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@param			data	body		request.CreateLocationRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/location [post]
func (h *locationHandler) CreateLocation(ctx *fiber.Ctx) error {
	req := new(request.CreateLocationRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.CreateLocation(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// UpdateLocation godoc
//
//	@Summary		Update Location
//	@Description	Location
//	@ID				update-location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			code	path		string							false	"Location Code"
//	@param			data	body		request.UpdateLocationRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/location/{code} [put]
func (h *locationHandler) UpdateLocation(ctx *fiber.Ctx) error {
	req := new(request.UpdateLocationRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID
	req.Code = ctx.Params("code")

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.UpdateLocation(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// DeleteLocationType godoc
//
//	@Summary		Delete Location
//	@Description	Location
//	@ID				delete-location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			code	path		string	true	"Code"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/location/{code} [delete]
func (h *locationHandler) DeleteLocation(ctx *fiber.Ctx) error {
	req := request.DeleteLocationRequest{}
	user := middleware.AuthUserData{}
	err := user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID
	req.Code = ctx.Params("code")

	h.service.DeleteLocation(ctx.Context(), req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}
