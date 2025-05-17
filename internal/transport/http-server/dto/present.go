package dto

type Present struct {
	Meeting    Meeting    `json:"Meeting"`
	Associated Associated `json:"Associated"`
	Date       string     `json:"date"`
	Present    bool       `json:"present"`
}
