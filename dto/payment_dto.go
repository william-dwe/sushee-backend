package dto

type PaymentOptionsResBody struct {
	PaymentOptions []PaymentOptionResBody `json:"payment_options"`
}

type PaymentOptionResBody struct {
	ID          int    `json:"id"`
	PaymentName string `json:"payment_name"`
}
