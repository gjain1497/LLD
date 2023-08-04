package main

type Payment struct{}

func (p *Payment) payBill(bill *Bill) {
	// do payment processing and update the bill status
	bill.isBillPaid = true
}
