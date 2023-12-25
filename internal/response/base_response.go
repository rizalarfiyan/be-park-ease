package response

import (
	"be-park-ease/internal/model"
	"be-park-ease/internal/request"

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

type BaseMetadataPagination struct {
	Total   int64 `json:"total"`
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	HasNext bool  `json:"has_next"`
}

type BaseResponsePagination[T any] struct {
	Content  []T                    `json:"content"`
	Metadata BaseMetadataPagination `json:"metadata"`
}

func WithPagination[T any](content model.ContentPagination[T], req request.BasePagination) BaseResponsePagination[T] {
	return BaseResponsePagination[T]{
		Content: content.Content,
		Metadata: BaseMetadataPagination{
			Total:   content.Count,
			Page:    req.Page,
			PerPage: req.Limit,
			HasNext: int64(req.Page*req.Limit) < content.Count,
		},
	}
}
