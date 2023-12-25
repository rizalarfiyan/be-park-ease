package response

import "github.com/jackc/pgx/v5/pgtype"

type VehicleType struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (vt *VehicleType) SetPrice(rawPrice pgtype.Numeric) {
	price, err := rawPrice.Float64Value()
	if rawPrice.Valid && err == nil {
		vt.Price = price.Float64
	}
}
