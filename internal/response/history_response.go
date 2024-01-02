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

type HistoryStatistic struct {
	RevenueTotal   float64                 `json:"revenue_total"`
	VehicleTotal   int                     `json:"vehicle_total"`
	CurrentVehicle int                     `json:"current_vehicle"`
	AvailableSpace int                     `json:"available_space"`
	ExitRevenue    float64                 `json:"exit_revenue"`
	ExitTotal      int                     `json:"exit_total"`
	FineRevenue    float64                 `json:"fine_revenue"`
	FineTotal      int                     `json:"fine_total"`
	Charts         []HistoryStatisticChart `json:"charts"`
}

type HistoryStatisticChart struct {
	Name    string  `json:"name"`
	Revenue float64 `json:"revenue"`
	Vehicle int     `json:"vehicle"`
}
