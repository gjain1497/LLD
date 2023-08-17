package main

import (
	pg "github.com/gjain1497/LLD/payment_gateway/payment_gateway"
	"log"
)

func payCreditCardHelper(p *pg.PaymentGateway) {
	//make payment credit card
	err1 := p.MakePaymentCreditCard(482920210, 123, "GIRISH JAIN")
	if err1 != nil {
		log.Printf("err1: %v\n", err1)
	}
	err2 := p.MakePaymentCreditCard(323233221, 234, "NIPUN JAIN")
	if err2 != nil {
		log.Printf("err2: %v\n", err2)
	}
	err3 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err3 != nil {
		log.Printf("err3: %v\n", err3)
	}
	err4 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err4 != nil {
		log.Printf("err4: %v\n", err4)
	}
	err5 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err5 != nil {
		log.Printf("err5: %v\n", err5)
	}
	err6 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err6 != nil {
		log.Printf("err6: %v\n", err6)
	}
	err7 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err7 != nil {
		log.Printf("err3: %v\n", err7)
	}
	err8 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err8 != nil {
		log.Printf("err8: %v\n", err8)
	}
	err9 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err9 != nil {
		log.Printf("err9: %v\n", err9)
	}
	err10 := p.MakePaymentCreditCard(4829222110, 149, "VANSH JAIN")
	if err10 != nil {
		log.Printf("err10: %v\n", err10)
	}

}

func payNetBankingHelper(p *pg.PaymentGateway) {
	err1 := p.MakePaymentNetBanking("123", "hello")
	if err1 != nil {
		log.Printf("err1: %v\n", err1)
	}
	err2 := p.MakePaymentNetBanking("123", "hello")
	if err2 != nil {
		log.Printf("err2: %v\n", err2)
	}
	err3 := p.MakePaymentNetBanking("123", "hello")
	if err3 != nil {
		log.Printf("err3: %v\n", err3)
	}
	err4 := p.MakePaymentNetBanking("123", "hello")
	if err4 != nil {
		log.Printf("err4: %v\n", err4)
	}
	err5 := p.MakePaymentNetBanking("123", "hello")
	if err5 != nil {
		log.Printf("err5: %v\n", err5)
	}
	err6 := p.MakePaymentNetBanking("123", "hello")
	if err6 != nil {
		log.Printf("err6: %v\n", err6)
	}
	err7 := p.MakePaymentNetBanking("123", "hello")
	if err7 != nil {
		log.Printf("err7: %v\n", err7)
	}
	err8 := p.MakePaymentNetBanking("123", "hello")
	if err8 != nil {
		log.Printf("err8: %v\n", err8)
	}
	err9 := p.MakePaymentNetBanking("123", "hello")
	if err9 != nil {
		log.Printf("err9: %v\n", err9)
	}
	err10 := p.MakePaymentNetBanking("123", "hello")
	if err1 != nil {
		log.Printf("err10: %v\n", err10)
	}
}

func addPaymentGateway() *pg.PaymentGateway {
	p, err := pg.NewPaymentGateway(
		pg.Options{
			Name: "RazorPay",
		})
	if err != nil {
		log.Printf("err1: %v\n", err)
	}
	log.Printf("p: %+v\n", p)

	//add client
	p.AddClient("Kite")
	p.AddClient("Grow")
	p.AddPaymentMode("Neft")
	p.AddPaymentMode("UPI")
	p.AddBank("IDBI", &pg.BankIDBI{}, 3)
	p.AddBank("ICICI", &pg.BankICICI{}, 7)
	p.AddBank("HDFC", &pg.BankHDFC{}, 3)
	p.AddBank("IDFC", &pg.BankIDFC{}, 7)
	return p
}

func main() {

	//makePaymentGateway
	p := addPaymentGateway()
	p.AddPaymentMode("CREDIT CARD")
	p.AddPaymentMode("DEBIT CARD")

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

	payCreditCardHelper(p)

	payNetBankingHelper(p)

	//make payment card
	//p.MakePaymentNetBanking("girishjain14", "Gjain@123")

	//show active distribution percentage of router
	errRouting := p.ShowRouterPercentage()
	if errRouting != nil {
		log.Printf("errRouting: %v\n", errRouting)
	}
}
