package handler

import (
	"be-park-ease/config"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
}

type authHandler struct {
	conf    *config.Config
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		conf:    config.Get(),
		service: service,
	}
}

// Auth Login godoc
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
	if err != nil {
		return err
	}

	// exception.ValidateStruct(*req, false)

	res := h.service.Login(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}
