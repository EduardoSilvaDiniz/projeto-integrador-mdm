package domain

type Present struct {
	Id         int        `json:"id"`
	Meeting    Meeting    `json:"Meeting"`
	Associated Associated `json:"Associated"`
	Date       string     `json:"date"`
	Present    bool       `json:"present"`
}
