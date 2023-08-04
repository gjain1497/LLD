package main

type Bill struct {
	reservation     *Reservation
	totalBillAmount float64
	isBillPaid      bool
}

func NewBill(reservation *Reservation) *Bill {
	bill := &Bill{
		reservation:     reservation,
		totalBillAmount: computeBillAmount(), // Calling the computeBillAmount function to calculate the bill amount
		isBillPaid:      false,
	}
	return bill
}

func computeBillAmount() float64 {
	return 100.0
}
