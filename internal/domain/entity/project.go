package entity

import "github.com/google/uuid"

type Project struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name,omitempty"`
	PaymentAmount float32   `json:"payment_amount,omitempty"`
}
