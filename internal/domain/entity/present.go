package entity

type Present struct {
	ID         int64      `json:"id"`
	Meeting    Meeting    `json:"meeting"`
	Associated Associated `json:"associated"`
	Date       string     `json:"date,omitempty"`
	Present    bool       `json:"present,omitempty"`
}
