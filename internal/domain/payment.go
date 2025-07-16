package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Payment struct {
	NumberCard  int64     `json:"number_card"`
	RefMonth    string    `json:"ref_month"`
	PaymentDate time.Time `json:"payment_date"`
}

func (a Payment) ToCreateParams() db.CreatePaymentParams {
	slog.Debug("chamada de função ToCreateParams em Payment")
	return db.CreatePaymentParams{
		NumberCard:  a.NumberCard,
		RefMonth:    a.RefMonth,
		PaymentDate: a.PaymentDate,
	}
}

func (a Payment) ToUpdateParams() db.UpdatePaymentParams {
	slog.Debug("chamada de função ToUpdateParams em Payment")
	return db.UpdatePaymentParams{
		RefMonth:    a.RefMonth,
		PaymentDate: a.PaymentDate,
	}
}
