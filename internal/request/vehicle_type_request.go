package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateVehicleTypeRequest struct {
	Code   string  `json:"code" example:"K001"`
	Name   string  `json:"name" example:"Bicycle"`
	Price  float64 `json:"price" example:"2000"`
	UserId int32   `json:"-"`
}

func (req CreateVehicleTypeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(2, 16)),
		validation.Field(&req.Name, validation.Required, validation.Length(5, 255)),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0), validation.Max(99999999999999999999.99)),
	)
}

type UpdateVehicleTypeRequest struct {
	Name   string  `json:"name" example:"Bicycle"`
	Price  float64 `json:"price" example:"2000"`
	Code   string  `json:"-"`
	UserId int32   `json:"-"`
}

func (req UpdateVehicleTypeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(2, 16).Error("Vehicle Type not found")),
		validation.Field(&req.Name, validation.Required, validation.Length(5, 255)),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0), validation.Max(99999999999999999999.99)),
	)
}

type DeleteVehicleTypeRequest struct {
	Code   string
	UserId int32
}
