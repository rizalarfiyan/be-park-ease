package handler

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type SettingHandler interface {
	GetAllSetting(ctx *fiber.Ctx) error
	CreateOrUpdateSetting(ctx *fiber.Ctx) error
}

type settingHandler struct {
	conf      *config.Config
	service   service.SettingService
	exception exception.Exception
}

func NewSettingHandler(service service.SettingService) SettingHandler {
	return &settingHandler{
		conf:      config.Get(),
		service:   service,
		exception: exception.NewException("setting-handler"),
	}
}

// GetAllSetting godoc
//
//	@Summary		Get All Setting based on parameter
//	@Description	Get All Setting
//	@ID				get-all-setting
//	@Tags			setting
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Success		200	{object}	response.BaseResponse{data=response.SettingResponse}
//	@Failure		500	{object}	response.BaseResponse
//	@Router			/setting [get]
func (h *settingHandler) GetAllSetting(ctx *fiber.Ctx) error {
	res := h.service.GetAllSetting(ctx.Context())
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data:    res,
	})
}

// CreateOrUpdateSetting godoc
//
//	@Summary		Post Create Or Update based on parameter
//	@Description	Create Or Update Setting
//	@ID				post-create-or-update-setting
//	@Tags			setting
//	@Accept			json
//	@Produce		json
//	@Security		AccessToken
//	@Param			data	body		request.CreateOrUpdateSettingRequest	true	"Data"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/setting [post]
func (h *settingHandler) CreateOrUpdateSetting(ctx *fiber.Ctx) error {
	req := new(request.CreateOrUpdateSettingRequest)
	err := ctx.BodyParser(req)
	h.exception.IsBadRequestErr(err, "Invalid request body", false)

	err = req.Validate()
	h.exception.IsErrValidation(err, true)

	h.service.CreateOrUpdateSetting(ctx.Context(), *req)
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
	})
}
