package domain

type Payment struct {
	Id             int    `json:"id"`
	Associated     int    `json:"associated"`
	Date           string `json:"date"`
	MonthReference string `json:"monthReference"`
}
