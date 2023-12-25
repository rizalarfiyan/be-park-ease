package handler

import (
	"be-park-ease/config"
	"be-park-ease/internal/response"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BaseHandler interface {
	Home(ctx *fiber.Ctx) error
}

type baseHandler struct {
	conf *config.Config
}

func NewBaseHandler() BaseHandler {
	return &baseHandler{
		conf: config.Get(),
	}
}

// Home godoc
//
//	@Summary		Get Base Home based on parameter
//	@Description	Base Home
//	@ID				get-base-home
//	@Tags			home
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.BaseResponse
//	@Failure		500	{object}	response.BaseResponse
//	@Router			/ [get]
func (h *baseHandler) Home(ctx *fiber.Ctx) error {
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data: fiber.Map{
			"title": fmt.Sprintf("%s API", h.conf.Name),
			"source_code": fiber.Map{
				"backend": "https://github.com/rizalarfiyan/be-park-ease",
				"winform": "https://github.com/rizalarfiyan/park-ease",
			},
			"author": []fiber.Map{
				{
					"name":   "Muhamad Rizal Arfiyan",
					"nim":    "22.11.5227",
					"github": "https://github.com/rizalarfiyan",
				},
				{
					"name":   "Ahmad Mufied Nugroho",
					"nim":    "22.11.5219",
					"github": "https://github.com/ahmad-mufied",
				},
				{
					"name":   "Damar Galih",
					"nim":    "22.11.52097",
					"github": "https://github.com/damar-glh",
				},
				{
					"name":   "Gilang Nur Hidayat",
					"nim":    "22.11.5196",
					"github": "https://github.com/glngnoor",
				},
				{
					"name":   "Wisnu Kusuma Dewa",
					"nim":    "22.11.5218",
					"github": "https://github.com/wisnu-kusuma-dewa",
				},
			},
		},
	})
}
