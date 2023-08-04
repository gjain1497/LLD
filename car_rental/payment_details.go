package main

import "time"

type PaymentMode int

const (
// Define PaymentMode constants here
)

type PaymentDetails struct {
	paymentId     int
	amountPaid    int
	dateOfPayment time.Time
	isRefundable  bool
	paymentMode   PaymentMode
}
