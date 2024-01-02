package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateOrUpdateSettingRequest struct {
	FineTicketCalculation int `json:"fine_ticket_calculation"`
	NextHourCalculation   int `json:"next_hour_calculation"`
	MaxCapacity           int `json:"max_capacity"`
}

func (req CreateOrUpdateSettingRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FineTicketCalculation, validation.Required),
		validation.Field(&req.NextHourCalculation, validation.Required),
		validation.Field(&req.MaxCapacity, validation.Required),
	)
}
