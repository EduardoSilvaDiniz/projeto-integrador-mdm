package domain

type Project struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	PaymentAmount float32 `json:"paymentAmount"`
}
