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
