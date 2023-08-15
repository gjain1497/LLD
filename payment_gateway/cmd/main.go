package main

import (
	"fmt"
	paymentgateway "github.com/gjain1497/LLD/payment_gateway/payment_gateway"
	"log"
)

func addPaymentGateway() *paymentgateway.PaymentGateway {
	p, err := paymentgateway.NewPaymentGateway(
		paymentgateway.Options{
			Name: "RazorPay",
		})
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}
	log.Printf("p: %+v\n", p)

	//add client
	p.AddClient("Kite")
	p.AddClient("Grow")
	p.AddPaymentMode("Neft")
	p.AddPaymentMode("UPI")
	p.AddBank(paymentgateway.NewBank("IDBI", "Chandigarh"))
	p.AddBank(paymentgateway.NewBank("ICICI", "Noida"))

	return p
}

func main() {

	//makePaymentGateway
	p := addPaymentGateway()

	//hasClient
	hasClient, err := p.HasClient("Kite")
	log.Printf("err1: %v\n", err)
	if hasClient {
		log.Println("Client Present")
	} else {
		log.Println("Client Absent")
	}

	//list all payment modes for a particular client
	p.ListPaymentModes()

	//show active distribution percentage of router
	p.ShowRouterPercentage()

	//make payment credit card
	p.MakePaymentCreditCard(482920210, 123, "GIRISH JAIN")

	//make payment card
	p.MakePaymentNetBanking("girishjain14", "Gjain@123")

}
