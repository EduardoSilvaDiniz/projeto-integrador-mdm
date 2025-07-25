package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Payment struct {
	ID          int64     `json:"id"`
	NumberCard  int64     `json:"number_card"  validate:"required,gt=0"`
	RefMonth    string    `json:"ref_month"    validate:"required"`
	PaymentDate time.Time `json:"payment_date" validate:"required"`
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
