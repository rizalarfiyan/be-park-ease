package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateOrUpdateSettingRequest struct {
	FineTicketCalculation int `json:"fine_ticket_calculation"`
	NextHourCalculation   int `json:"next_hour_calculation"`
}

func (req CreateOrUpdateSettingRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FineTicketCalculation, validation.Required),
		validation.Field(&req.NextHourCalculation, validation.Required),
	)
}
