package response

import (
	"be-park-ease/constants"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
)

type SettingResponse struct {
	FineTicketCalculation            int    `json:"fine_ticket_calculation"`
	FineTicketCalculationDescription string `json:"fine_ticket_calculation_description"`
	NextHourCalculation              int    `json:"next_hour_calculation"`
	NextHourCalculationDescription   string `json:"next_hour_calculation_description"`
	MaxCapacity                      int    `json:"max_capacity"`
	MaxCapacityDescription           string `json:"max_capacity_description"`
}

func (r *SettingResponse) FromDB(settings []sql.Setting) {
	for _, setting := range settings {
		switch setting.Key {
		case constants.SettingFineTicketCalculation:
			r.FineTicketCalculation, _ = utils.StrToInt(setting.Value)
			r.FineTicketCalculationDescription = setting.Description
		case constants.SettingNextHourCalculation:
			r.NextHourCalculation, _ = utils.StrToInt(setting.Value)
			r.NextHourCalculationDescription = setting.Description
		case constants.SettingMaxCapacity:
			r.MaxCapacity, _ = utils.StrToInt(setting.Value)
			r.MaxCapacityDescription = setting.Description
		}
	}
}
