package domain

import (
	"projeto-integrador-mdm/internal/database"
	"time"
)

type Payment struct {
	NumberCard  int64     `json:"number_card"`
	RefMonth    string    `json:"ref_month"`
	PaymentDate time.Time `json:"payment_date"`
}

func (a Payment) ToCreateParams() database.CreatePaymentParams {
	return database.CreatePaymentParams{
		NumberCard:  a.NumberCard,
		RefMonth:    a.RefMonth,
		PaymentDate: a.PaymentDate,
	}
}
