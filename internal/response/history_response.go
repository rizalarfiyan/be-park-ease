package response

import "be-park-ease/constants"

type EntryHistory struct {
	ID              string                `json:"id"`
	LocationCode    string                `json:"location_code"`
	VehicleTypeCode string                `json:"vehicle_type_code"`
	VehicleNumber   string                `json:"vehicle_number"`
	Type            constants.HistoryType `json:"type"`
	Date            string                `json:"date"`
}
