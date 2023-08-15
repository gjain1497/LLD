package payment_gateway

import (
	"fmt"
	"log"
	"sync"
)

//
type Options struct {
	Name string
}
type PaymentGateway struct {
	name         string
	clients      map[string]bool
	paymentModes map[string]bool
	banks        map[string]*Bank
	router       *Router
	mutex        sync.Mutex
}

func NewPaymentGateway(o Options) (*PaymentGateway, error) {

	p := &PaymentGateway{
		name:         o.Name,
		clients:      map[string]bool{},
		paymentModes: map[string]bool{},
		banks:        map[string]*Bank{},
		router:       &Router{},
	}
	return p, nil

}

func (p *PaymentGateway) AddClient(client string) {
	p.clients[client] = true
}

func (p *PaymentGateway) RemoveClient(client string) {
	delete(p.clients, client)
}
func (p *PaymentGateway) HasClient(client string) (bool, error) {
	_, exists := p.clients[client]
	if exists {
		fmt.Printf("Client '%s' exists", client)
		return true, nil
	}
	fmt.Printf("Client '%s' does not exist", client)
	return false, nil
}

func (p *PaymentGateway) AddPaymentMode(paymentMode string) {
	p.paymentModes[paymentMode] = true
}

func (p *PaymentGateway) ListPaymentModes() {
	for mode, _ := range p.paymentModes {
		log.Print(mode, " ")
	}
	log.Println()
}

func (p *PaymentGateway) RemovePaymentMode(paymentMode string) {
	delete(p.paymentModes, paymentMode)
}

func (p *PaymentGateway) AddBank(bank *Bank) {
	p.banks[bank.name] = bank
}

func (p *PaymentGateway) RemoveBank(bank *Bank) {
	delete(p.banks, bank.name)
}

func (p *PaymentGateway) MakePaymentCreditCard(cardNumber int, cvv int, cardName string) error {
	payment, err := p.router.MakePaymentCreditCard(cardNumber, cvv, cardName)
	if err != nil {
		fmt.Printf("err completing payment: %v\n", err)
		return err
	}
	log.Printf("payment: %+v\n", payment)
	return nil
}

func (p *PaymentGateway) MakePaymentNetBanking(userID string, password string) error {
	payment, err := p.router.MakePaymentNetBanking(userID, password)
	if err != nil {
		fmt.Printf("err completing payment: %v\n", err)
		return err
	}
	log.Printf("payment: %+v\n", payment)
	return nil
}

func (p *PaymentGateway) ShowRouterPercentage() {

}
