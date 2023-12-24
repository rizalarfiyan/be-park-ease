package request

import "be-park-ease/constants"

type AllHistoryRequest struct {
	BasePagination
	HistoryType constants.HistoryType
	VehicleType string
	Location    string
}

func (r *AllHistoryRequest) Normalize() {
	if !r.HistoryType.IsValid() && r.HistoryType != "" {
		r.HistoryType = constants.HistoryTypeEntry
	}

	r.BasePagination.Normalize()
}
