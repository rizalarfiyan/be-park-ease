package response

import "time"

type VehicleType struct {
	Code  string    `json:"code"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
	Date  time.Time `json:"date"`
}
