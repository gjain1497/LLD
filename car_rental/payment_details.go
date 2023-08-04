package main

import "time"

type PaymentMode int

const (
	CASH PaymentMode = iota
	ONLINE
)

type PaymentDetails struct {
	paymentId     int
	amountPaid    int
	dateOfPayment time.Time
	isRefundable  bool
	paymentMode   PaymentMode
}
