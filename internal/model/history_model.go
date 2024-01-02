package model

import (
    "be-park-ease/constants"
    "time"
)

type AllHistoryStatistic struct {
	StartDate     time.Time
	EndDate       time.Time
	TimeFrequency constants.FilterTimeFrequency
}
