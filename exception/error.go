package exception

import (
	"be-park-ease/internal/response"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := http.StatusText(code)

	var fiberError *fiber.Error
	var data interface{} = err.Error()
	if errors.As(err, &fiberError) {
		code = fiberError.Code
		message = http.StatusText(code)
		if !strings.EqualFold(err.Error(), message) {
			message = err.Error()
		}
		data = nil
	}

	var httpError *response.BaseResponse
	if errors.As(err, &httpError) {
		code = httpError.Code
		data = httpError.Data
		if httpError.Message != "" {
			message = httpError.Message
		} else {
			message = http.StatusText(code)
		}
	}

	return ctx.Status(code).JSON(response.BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
