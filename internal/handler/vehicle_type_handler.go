package handler

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type VehicleTypeHandler interface {
	AllVehicleType(ctx *fiber.Ctx) error
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
//	@Summary		Get All VehicleType based on parameter
//	@Description	All VehicleType
//	@ID				get-all-vehicleType
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
