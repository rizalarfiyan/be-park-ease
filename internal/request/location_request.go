package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type GetAllLocationRequest struct {
	BasePagination
	IsExit *bool
}

type CreateLocationRequest struct {
	Code   string `json:"code" example:"D1"`
	Name   string `json:"name" example:"DOM 1"`
	IsExit bool   `json:"is_exit" example:"false"`
	UserId int32  `json:"-"`
}

func (req CreateLocationRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(2, 16)),
		validation.Field(&req.Name, validation.Required, validation.Length(5, 255)),
	)
}

type UpdateLocationRequest struct {
	Name   string `json:"name" example:"DOM 1"`
	IsExit bool   `json:"is_exit" example:"false"`
	Code   string `json:"-"`
	UserId int32  `json:"-"`
}

func (req UpdateLocationRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(2, 16).Error("Location not found")),
		validation.Field(&req.Name, validation.Required, validation.Length(5, 255)),
		validation.Field(&req.UserId, validation.Required),
	)
}

type DeleteLocationRequest struct {
	Code   string
	UserId int32
}
