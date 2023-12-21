package response

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Code    int         `json:"code" example:"999"`
	Message string      `json:"message" example:"Message!"`
	Data    interface{} `json:"data"`
}

func (res *BaseResponse) Error() string {
	return res.Message
}

func NewError(code int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code: code,
		Data: data,
	}
}

func NewErrorMessage(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		code,
		message,
		data,
	}
}

func New(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	return ctx.Status(code).JSON(&BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
