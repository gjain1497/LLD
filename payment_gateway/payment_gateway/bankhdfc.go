package payment_gateway

type BankHDFC struct{}

func (b *BankHDFC) CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	return NewPayment("COMPLETED", "CREDITCARD"), nil
}

func (b *BankHDFC) CompletePaymentNetBanking(userID string, password string) (*Payment, error) {
	return NewPayment("COMPLETED", "NETBANKING"), nil
}
