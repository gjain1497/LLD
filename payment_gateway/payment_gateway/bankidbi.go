package payment_gateway

type BankIDBI struct{}

func (b *BankIDBI) CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	return NewPayment("COMPLETED", "CREDITCARD"), nil
}

func (b *BankIDBI) CompletePaymentNetBanking(userID string, password string) (*Payment, error) {
	return NewPayment("COMPLETED", "NETBANKING"), nil
}
