package handler

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"be-park-ease/internal/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler interface {
	AllUser(ctx *fiber.Ctx) error
}

type userHandler struct {
	conf      *config.Config
	service   service.UserService
	exception exception.Exception
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("user-handler"),
	}
}

// AllUser godoc
//
//	@Summary		Get All User based on parameter
//	@Description	All User
//	@ID				get-all-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			page		query		int		false	"Page"	default(1)
//	@Param			limit		query		int		false	"Limit"	default(10)
//	@Param			search		query		string	false	"Search"
//	@Param			order_by	query		string	false	"Order by"	Enums(name,username,role,status)
//	@Param			order		query		string	false	"Order"		Enums(asc, desc)
//	@Param			role		query		string	false	"Role"		Enums(admin,karyawan)
//	@Param			status		query		string	false	"Status"	Enums(active,banned)
//	@Success		200			{object}	response.BaseResponse{data=response.BaseResponsePagination[response.User]}
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/user [get]
func (h *userHandler) AllUser(ctx *fiber.Ctx) error {
	req := request.AllUserRequest{
		BasePagination: request.BasePagination{
			Page:    ctx.QueryInt("page", 1),
			Limit:   ctx.QueryInt("limit", constants.DefaultPageLimit),
			Search:  ctx.Query("search"),
			OrderBy: ctx.Query("order_by"),
			Order:   ctx.Query("order"),
		},
		Role:   sql.UserRole(ctx.Query("role")),
		Status: sql.UserStatus(ctx.Query("status")),
	}

	fieldOrder := map[string]string{
		"name":     "name",
		"username": "username",
		"role":     "role",
		"status":   "status",
	}

	req.ValidateAndUpdateOrderBy(fieldOrder)
	req.Normalize()

	res := h.service.AllUser(ctx.Context(), req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}
