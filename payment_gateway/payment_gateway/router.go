package payment_gateway

type Router struct {
	id    int
	banks map[string]*Bank
}

func (r *Router) MakePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	bankIDBI := r.banks["IDBI"]
	payment, err := bankIDBI.CompletePaymentCreditCard(cardNumber, cvv, cardName)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *Router) MakePaymentNetBanking(userID string, password string) (*Payment, error) {
	bankIDBI := r.banks["ICICI"]
	payment, err := bankIDBI.CompletePaymentNetBanking(userID, password)
	if err != nil {
		return payment, err
	}
	return payment, nil
}
