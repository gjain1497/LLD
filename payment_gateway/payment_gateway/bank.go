package payment_gateway

type Bank struct {
	name   string
	branch string
}

func NewBank(name string, branch string) *Bank {
	return &Bank{name: name, branch: branch}
}

func (b *Bank) CompletePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	return NewPayment("COMPLETED", "CREDITCARD"), nil
}

func (b *Bank) CompletePaymentNetBanking(userID string, password string) (*Payment, error) {
	return NewPayment("COMPLETED", "NETBANKING"), nil
}
