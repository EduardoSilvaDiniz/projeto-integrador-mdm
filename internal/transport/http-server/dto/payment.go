package dto

type Payment struct {
	Associated     int    `json:"associated"`
	Date           string `json:"date"`
	MonthReference string `json:"monthReference"`
}
