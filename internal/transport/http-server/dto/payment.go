package dto

type Payment struct {
	Associated     int    `json:"associated,omitempty"`
	Date           string `json:"date,omitempty"`
	MonthReference string `json:"month_reference,omitempty"`
}
