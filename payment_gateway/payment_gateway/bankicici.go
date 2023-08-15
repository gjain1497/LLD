package payment_gateway

type BankICICI struct{}

func (b *BankICICI) CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	return NewPayment("COMPLETED", "CREDITCARD"), nil
}

func (b *BankICICI) CompletePaymentNetBanking(userID string, password string) (*Payment, error) {
	return NewPayment("COMPLETED", "NETBANKING"), nil
}
