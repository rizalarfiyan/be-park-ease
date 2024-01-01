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
	LocationCode    string `json:"location_code" example:"DOM001"`
	VehicleTypeCode string `json:"vehicle_type_code" example:"K001"`
	VehicleNumber   string `json:"vehicle_number" example:"AB2342NW"`
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
	EntryHistoryId string `json:"entry_history_id" example:"H251845879AA5F13"`
	IsFine         bool   `json:"is_fine" example:"false"`
}

func (req CalculatePriceHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntryHistoryId, validation.Required, validation.Length(14, 16).Error("Entry History not found")),
	)
}

type CreateExitHistoryRequest struct {
	EntryHistoryId string  `json:"entry_history_id" example:"H251845879AA5F13"`
	LocationCode   string  `json:"location_code" example:"DOM002"`
	Price          float64 `json:"price" example:"32500"`
	UserId         int32   `json:"-"`
}

func (req CreateExitHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntryHistoryId, validation.Required, validation.Length(14, 16).Error("Entry History not found")),
		validation.Field(&req.LocationCode, validation.Required, validation.Length(2, 16).Error("Location Type not found")),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0), validation.Max(99999999999999999999.99)),
	)
}

type CreateFineHistoryRequest struct {
	EntryHistoryId  string  `json:"entry_history_id" example:"H251845879AA5F13"`
	LocationCode    string  `json:"location_code" example:"DOM002"`
	Price           float64 `json:"price" example:"32500"`
	Identity        string  `json:"identity" example:"1234567890123456"`
	VehicleIdentity string  `json:"vehicle_identity" example:"AB2342NW"`
	Name            string  `json:"name" example:"Paijo Royo Royo"`
	Address         string  `json:"address" example:"Jl. Ring Road Utara, Ngringin, Condongcatur, Kec. Depok, Kabupaten Sleman"`
	Description     string  `json:"description" example:"My ticket is lost"`
	UserId          int32   `json:"-"`
}

func (req CreateFineHistoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntryHistoryId, validation.Required, validation.Length(14, 16).Error("Entry History not found")),
		validation.Field(&req.LocationCode, validation.Required, validation.Length(2, 16).Error("Location Type not found")),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0), validation.Max(99999999999999999999.99)),
		validation.Field(&req.Identity, validation.Required, validation.Length(5, 32)),
		validation.Field(&req.VehicleIdentity, validation.Required, validation.Length(5, 32)),
		validation.Field(&req.Name, validation.Required, validation.Length(5, 100)),
		validation.Field(&req.Address, validation.Required, validation.Length(10, 255)),
		validation.Field(&req.Description, validation.Required, validation.Length(10, 255)),
	)
}
