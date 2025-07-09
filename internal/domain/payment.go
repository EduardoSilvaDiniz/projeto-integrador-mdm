package domain

type Payment struct {
	ID             int64  `json:"id"`
	Associated     int    `json:"associated,omitempty"`
	Date           string `json:"date,omitempty"`
	MonthReference string `json:"month_reference,omitempty"`
}
