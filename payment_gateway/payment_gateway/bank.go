package payment_gateway

type Bank interface {
	CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error)
	CompletePaymentNetBanking(userID string, password string) (*Payment, error)
}
