package request

import (
	"be-park-ease/constants"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AllHistoryRequest struct {
	BasePagination
	HistoryType constants.HistoryType
	VehicleType string
	Location    string
}

func (r *AllHistoryRequest) Normalize() {
	if r.HistoryType != "" && !r.HistoryType.IsValid() {
		r.HistoryType = constants.HistoryTypeEntry
	}

	r.BasePagination.Normalize()
}

type CreateEntryHistoryRequest struct {
	LocationCode    string `json:"location_code"`
	VehicleTypeCode string `json:"vehicle_type_code"`
	VehicleNumber   string `json:"vehicle_number"`
	UserId          int32  `json:"-"`
}

func (req CreateEntryHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.LocationCode, validation.Required, validation.Length(2, 16).Error("Location Type not found")),
		validation.Field(&req.VehicleTypeCode, validation.Required, validation.Length(2, 16).Error("Vehicle Type not found")),
		validation.Field(&req.VehicleNumber, validation.Required, validation.Length(3, 16), constants.ValidationVehicleNumber),
	)
}

type CalculatePriceHistoryRequest struct {
	VehicleTypeCode string `json:"vehicle_type_code"`
	VehicleNumber   string `json:"vehicle_number"`
	IsFine          bool   `json:"is_fine"`
}

func (req CalculatePriceHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.VehicleTypeCode, validation.Required, validation.Length(2, 16).Error("Vehicle Type not found")),
		validation.Field(&req.VehicleNumber, validation.Required, validation.Length(3, 16), constants.ValidationVehicleNumber),
	)
}

type CreateExitHistoryRequest struct {
	EntryHistoryId string  `json:"entry_history_id"`
	LocationCode   string  `json:"location_code"`
	Price          float64 `json:"price"`
	UserId         int32   `json:"-"`
}

func (req CreateExitHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntryHistoryId, validation.Required, validation.Length(14, 16).Error("Entry History not found")),
		validation.Field(&req.LocationCode, validation.Required, validation.Length(2, 16).Error("Location Type not found")),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0), validation.Max(99999999999999999999.99)),
	)
}
