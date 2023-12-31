package handler

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"be-park-ease/internal/sql"
	"be-park-ease/middleware"
	"be-park-ease/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	AllUser(ctx *fiber.Ctx) error
	UserById(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
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
//	@Param			order_by	query		string	false	"Order by"	Enums(name,username,role,status,date)
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
		"date":     "created_at",
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

// UserById godoc
//
//	@Summary		Get User By ID based on parameter
//	@Description	User By ID
//	@ID				get-user-by-id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			id	path		int	false	"User ID"
//	@Success		200	{object}	response.BaseResponse{data=response.User}
//	@Failure		500	{object}	response.BaseResponse
//	@Router			/user/{id} [get]
func (h *userHandler) UserById(ctx *fiber.Ctx) error {
	userIdStr := ctx.Params("id")
	userid, err := utils.StrToInt(userIdStr)
	h.exception.IsBadRequestErr(err, "Invalid user id", false)

	res := h.service.UserById(ctx.Context(), int32(userid))
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// CreateUser godoc
//
//	@Summary		Post Create User based on parameter
//	@Description	Create User
//	@ID				post-create-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.CreateUserRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/user [post]
func (h *userHandler) CreateUser(ctx *fiber.Ctx) error {
	req := new(request.CreateUserRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.CreateUser(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// UpdateUser godoc
//
//	@Summary		Post Update User based on parameter
//	@Description	Update User
//	@ID				put-update-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			id		path		int							false	"User ID"
//	@Param			data	body		request.UpdateUserRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/user/{id} [put]
func (h *userHandler) UpdateUser(ctx *fiber.Ctx) error {
	req := new(request.UpdateUserRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	userIdStr := ctx.Params("id")
	userid, err := utils.StrToInt(userIdStr)
	h.exception.IsBadRequestErr(err, "Invalid user id", false)
	req.UserId = int32(userid)

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.UpdateUser(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}

// ChangePassword godoc
//
//	@Summary		Post Change Password User based on parameter
//	@Description	Post Change Password
//	@ID				post-change-password-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.ChangePasswordRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/user/change-password [post]
func (h *userHandler) ChangePassword(ctx *fiber.Ctx) error {
	req := new(request.ChangePasswordRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	user := middleware.AuthUserData{}
	err = user.Get(ctx)
	h.exception.PanicIfError(err, false)
	req.UserId = user.ID

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	h.service.ChangePassword(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}
