package response

type Location struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	IsExit bool   `json:"is_exit"`
}
