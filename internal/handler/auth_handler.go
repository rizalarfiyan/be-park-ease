package handler

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"be-park-ease/middleware"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}

type authHandler struct {
	conf      *config.Config
	service   service.AuthService
	exception exception.Exception
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("auth-handler"),
	}
}

// Login godoc
//
//	@Summary		Post Auth Login based on parameter
//	@Description	Auth Login
//	@ID				post-auth-login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			data	body		request.AuthLoginRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/auth/login [post]
func (h *authHandler) Login(ctx *fiber.Ctx) error {
	req := new(request.AuthLoginRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	err = req.Validate()
	h.exception.IsErrValidation(err, false)

	res := h.service.Login(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// Me godoc
//
//	@Summary		Get Auth Me based on parameter
//	@Description	Auth Me
//	@ID				get-auth-me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Success		200	{object}	response.BaseResponse{data=middleware.AuthUserData}
//	@Failure		500	{object}	response.BaseResponse
//	@Router			/auth/me [get]
func (h *authHandler) Me(ctx *fiber.Ctx) error {
	user := middleware.AuthUserData{}
	err := user.Get(ctx)
	h.exception.PanicIfError(err, false)

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    user,
	})
}
