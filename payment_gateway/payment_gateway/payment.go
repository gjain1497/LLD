package payment_gateway

import "github.com/google/uuid"

type Payment struct {
	paymentId     string
	paymentStatus string
	paymentMode   string
}

func NewPayment(paymentStatus string, paymentMode string) *Payment {
	return &Payment{
		paymentId:     uuid.NewString(),
		paymentStatus: paymentStatus,
		paymentMode:   paymentMode}
}

//addPaymentMode
