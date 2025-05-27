package dto

type Project struct {
	Name          string  `json:"name,omitempty"`
	PaymentAmount float32 `json:"payment_amount,omitempty"`
}
