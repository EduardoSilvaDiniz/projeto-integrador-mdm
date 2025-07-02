package entity

import "github.com/google/uuid"

type Payment struct {
	ID             uuid.UUID `json:"id"`
	Associated     int       `json:"associated,omitempty"`
	Date           string    `json:"date,omitempty"`
	MonthReference string    `json:"month_reference,omitempty"`
}
