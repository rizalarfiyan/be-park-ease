package handler

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"be-park-ease/middleware"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type VehicleTypeHandler interface {
	AllVehicleType(ctx *fiber.Ctx) error
	VehicleTypeByCode(ctx *fiber.Ctx) error
	CreateVehicleType(ctx *fiber.Ctx) error
	UpdateVehicleType(ctx *fiber.Ctx) error
}

type vehicleTypeHandler struct {
	conf      *config.Config
	service   service.VehicleTypeService
	exception exception.Exception
}

func NewVehicleTypeHandler(service service.VehicleTypeService) VehicleTypeHandler {
	return &vehicleTypeHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("vehicleType-handler"),
	}
}

// AllVehicleType godoc
//
//	@Summary		Get All Vehicle Type based on parameter
//	@Description	All Vehicle Type
//	@ID				get-all-vehicle-type
//	@Tags			vehicle-type
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page		query		int		false	"Page"	default(1)
//	@Param			limit		query		int		false	"Limit"	default(10)
//	@Param			search		query		string	false	"Search"
//	@Param			order_by	query		string	false	"Order by"	Enums(code,name,price,date)
//	@Param			order		query		string	false	"Order"		Enums(asc, desc)
//	@Success		200			{object}	response.BaseResponse{data=response.BaseResponsePagination[response.VehicleType]}
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/vehicle_type [get]
func (h *vehicleTypeHandler) AllVehicleType(ctx *fiber.Ctx) error {
	req := request.BasePagination{
		Page:    ctx.QueryInt("page", 1),
		Limit:   ctx.QueryInt("limit", constants.DefaultPageLimit),
		Search:  ctx.Query("search"),
		OrderBy: ctx.Query("order_by"),
		Order:   ctx.Query("order"),
	}

	fieldOrder := map[string]string{
		"code":  "code",
		"name":  "name",
		"price": "price",
		"date":  "created_at",
	}

	req.ValidateAndUpdateOrderBy(fieldOrder)
	req.Normalize()

	res := h.service.AllVehicleType(ctx.Context(), req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// VehicleTypeByCode godoc
//
//	@Summary		Get Vehicle Type By Code based on parameter
//	@Description	Vehicle Type By Code
//	@ID				get-vehicle-type-by-code
//	@Tags			vehicle-type
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			code	path		string	false	"Vehicle Type Code"
//	@Success		200		{object}	response.BaseResponse{data=response.VehicleType}
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/vehicle_type/{code} [get]
func (h *vehicleTypeHandler) VehicleTypeByCode(ctx *fiber.Ctx) error {
	res := h.service.VehicleTypeById(ctx.Context(), ctx.Params("code"))
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// CreateVehicleType godoc
//
//	@Summary		Post Create Vehicle Type based on parameter
//	@Description	Create Vehicle Type
//	@ID				post-create-vehicle-type
//	@Tags			vehicle-type
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.CreateVehicleTypeRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/vehicle_type [post]
func (h *vehicleTypeHandler) CreateVehicleType(ctx *fiber.Ctx) error {
	req := new(request.CreateVehicleTypeRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID

	err = req.Validate()
	h.exception.IsErrValidation(err, true)

	h.service.CreateVehicleType(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// UpdateVehicleType godoc
//
//	@Summary		Put Update Vehicle Type based on parameter
//	@Description	Update Vehicle Type
//	@ID				put-update-vehicle-type
//	@Tags			vehicle-type
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			code	path		string								false	"Vehicle Type Code"
//	@Param			data	body		request.UpdateVehicleTypeRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/vehicle_type/{code} [put]
func (h *vehicleTypeHandler) UpdateVehicleType(ctx *fiber.Ctx) error {
	req := new(request.UpdateVehicleTypeRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID
	req.Code = ctx.Params("code")

	err = req.Validate()
	h.exception.IsErrValidation(err, true)

	h.service.UpdateVehicleType(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}
