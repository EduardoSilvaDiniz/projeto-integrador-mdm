package entity

import "github.com/google/uuid"

type Present struct {
	ID         uuid.UUID  `json:"id"`
	Meeting    Meeting    `json:"meeting,omitempty"`
	Associated Associated `json:"associated,omitempty"`
	Date       string     `json:"date,omitempty"`
	Present    bool       `json:"present,omitempty"`
}
