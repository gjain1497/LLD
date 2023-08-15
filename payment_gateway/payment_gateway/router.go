package payment_gateway

import (
	"log"
	"math/rand"
)

type Router struct {
	banks                 map[Bank]string
	countTraffic          map[string]int
	routingDataCreditCard []Bank
	routingDataNetBanking []Bank
}
type RoutingEntry struct {
	Bank       Bank
	Percentage float64
}

func NewRouter(id int, banks map[string]*Bank, countTraffic map[string]int) *Router {
	return &Router{
		banks:                 make(map[Bank]string),
		countTraffic:          make(map[string]int),
		routingDataCreditCard: make([]Bank, 0),
		routingDataNetBanking: make([]Bank, 0),
	}
}

func (r *Router) ShowRouterPercentage() error {
	log.Println("Count of payments done by IDBI Bank ", r.countTraffic["IDBI"])
	log.Println("Count of payments done by ICICI Bank ", r.countTraffic["ICICI"])
	log.Println("Count of payments done by HDFC Bank ", r.countTraffic["HDFC"])
	log.Println("Count of payments done by IDFC Bank ", r.countTraffic["IDFC"])
	return nil
}

func (r *Router) MakePaymentCreditCard(cardNumber int, cvv int, cardName string) (*Payment, error) {
	// Logic to determine which bank to use based on routing strategy
	selectedIndex := rand.Intn(len(r.routingDataCreditCard))
	selectedBank := r.routingDataCreditCard[selectedIndex]
	selectedBankName := r.banks[selectedBank]
	log.Println("Credit Card Payment done by ", selectedBankName)

	payment, err := selectedBank.CompletePaymentCreditCard(cardNumber, cvv, cardName)
	if err != nil {
		return payment, err
	}
	r.countTraffic[selectedBankName]++
	return payment, nil
}

func (r *Router) MakePaymentNetBanking(userID string, password string) (*Payment, error) {
	// Logic to determine which bank to use based on routing strategy
	selectedIndex := rand.Intn(len(r.routingDataNetBanking))
	selectedBank := r.routingDataNetBanking[selectedIndex]
	selectedBankName := r.banks[selectedBank]
	log.Println("NetBanking Payment done by ", selectedBankName)

	payment, err := selectedBank.CompletePaymentNetBanking(userID, password)
	if err != nil {
		return payment, err
	}
	r.countTraffic[selectedBankName]++
	return payment, nil
}
