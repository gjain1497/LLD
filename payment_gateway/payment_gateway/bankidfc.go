package payment_gateway

type BankIDFC struct{}

func (b *BankIDFC) CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	return NewPayment("COMPLETED", "CREDITCARD"), nil
}

func (b *BankIDFC) CompletePaymentNetBanking(userID string, password string) (*Payment, error) {
	return NewPayment("COMPLETED", "NETBANKING"), nil
}
