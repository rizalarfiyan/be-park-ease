package request

import (
	"be-park-ease/constants"
	"be-park-ease/utils"
	"html"
	"strings"
	"time"
)

type BasePagination struct {
	Page    int    `json:"page" field:"Page" validate:"omitempty,min=1" example:"1"`
	Limit   int    `json:"limit" field:"Limit" example:"10"`
	Search  string `json:"search" field:"Search"`
	OrderBy string `json:"order_by" field:"Order By"`
	Order   string `json:"order" field:"Order"`
}

func (bp *BasePagination) Normalize() {
	if bp.Search != "" {
		bp.Search = html.EscapeString(utils.Str(bp.Search))
	}
}

func (bp *BasePagination) ValidateAndUpdateOrderBy(data map[string]string) {
	keyOrderBy := strings.ToLower(bp.OrderBy)
	getOrderBy, isFoundOrderBy := data[keyOrderBy]
	if !isFoundOrderBy {
		bp.OrderBy = ""
		bp.Order = ""
		return
	}

	bp.OrderBy = getOrderBy
	orderMap := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	keyOrder := strings.ToLower(bp.Order)
	if _, ok := orderMap[keyOrder]; ok {
		bp.Order = keyOrder
	} else {
		bp.Order = "asc"
	}
}

type TimeFrequency struct {
	TimeFrequency constants.FilterTimeFrequency `json:"time_range"`
}

type TimeFrequencyBuilder struct {
	StartDate    time.Time
	EndDate      time.Time
	CallbackName func(time.Time) string
	CallbackDate func(time.Time) time.Time
}

func (tf *TimeFrequency) Normalize() {
	if !tf.TimeFrequency.IsValid() {
		tf.TimeFrequency = constants.FilterTimeFrequencyToday
	}
}

func (tf *TimeFrequency) BuildTimeFrequency() TimeFrequencyBuilder {
	tf.Normalize()
	res := TimeFrequencyBuilder{}
	now := time.Now()
	res.StartDate = utils.StartOfDay(now)
	switch tf.TimeFrequency {
	case constants.FilterTimeFrequencyWeek:
		res.StartDate = utils.StartOfWeek(now)
		res.EndDate = utils.EndOfWeek(now)
		res.CallbackName = func(val time.Time) string {
			return val.Format(time.DateOnly)
		}
		res.CallbackDate = func(val time.Time) time.Time {
			return val.AddDate(0, 0, 1)
		}
	case constants.FilterTimeFrequencyMonth:
		res.StartDate = utils.StartOfMonth(now)
		res.EndDate = utils.EndOfMonth(now)
		res.CallbackName = func(val time.Time) string {
			return val.Format(time.DateOnly)
		}
		res.CallbackDate = func(val time.Time) time.Time {
			return val.AddDate(0, 0, 1)
		}
	case constants.FilterTimeFrequencyQuarter:
		res.StartDate = utils.StartOfMonth(res.StartDate.AddDate(0, -6, 0))
		res.EndDate = utils.EndOfMonth(now)
		res.CallbackName = func(val time.Time) string {
			return val.Format("Jan 2006")
		}
		res.CallbackDate = func(val time.Time) time.Time {
			return val.AddDate(0, 1, 0)
		}
	case constants.FilterTimeFrequencyYear:
		res.StartDate = utils.StartOfMonth(res.StartDate.AddDate(-1, 0, 0))
		res.EndDate = utils.EndOfMonth(now)
		res.CallbackName = func(val time.Time) string {
			return val.Format("Jan 2006")
		}
		res.CallbackDate = func(val time.Time) time.Time {
			return val.AddDate(0, 1, 0)
		}
	default:
		res.EndDate = utils.EndOfDay(now)
		res.CallbackName = func(val time.Time) string {
			return val.Format(time.TimeOnly)
		}
		res.CallbackDate = func(val time.Time) time.Time {
			return val.Add(1 * time.Hour)
		}
	}

	return res
}
