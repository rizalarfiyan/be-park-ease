package constants

type HistoryType string

var (
	HistoryTypeEntry HistoryType = "entry"
	HistoryTypeExit  HistoryType = "exit"
	HistoryTypeFine  HistoryType = "fine"
)

func (h *HistoryType) IsValid() bool {
	switch *h {
	case HistoryTypeEntry, HistoryTypeExit, HistoryTypeFine:
		return true
	}
	return false
}

type FilterTimeFrequency string

const (
	FilterTimeFrequencyToday   FilterTimeFrequency = "today"
	FilterTimeFrequencyWeek    FilterTimeFrequency = "week"
	FilterTimeFrequencyMonth   FilterTimeFrequency = "month"
	FilterTimeFrequencyQuarter FilterTimeFrequency = "quarter"
	FilterTimeFrequencyYear    FilterTimeFrequency = "year"
)

func (v FilterTimeFrequency) IsValid() bool {
	switch v {
	case FilterTimeFrequencyToday, FilterTimeFrequencyWeek, FilterTimeFrequencyMonth, FilterTimeFrequencyQuarter, FilterTimeFrequencyYear:
		return true
	}

	return false
}
