package dto

type Present struct {
	Meeting    Meeting    `json:"meeting,omitempty"`
	Associated Associated `json:"associated,omitempty"`
	Date       string     `json:"date,omitempty"`
	Present    bool       `json:"present,omitempty"`
}
