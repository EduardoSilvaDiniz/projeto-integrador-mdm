package domain

type Meeting struct {
	ID        int64  `json:"id"`
	Date      string `json:"date,omitempty"`
	Time      string `json:"time,omitempty"`
}
