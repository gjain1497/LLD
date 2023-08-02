package main

// Define a custom struct to represent the Coin enum-like behavior.
type Coin struct {
	name  string
	value int
}

// Declare constants representing each coin.
var (
	PENNY   = Coin{"PENNY", 1}
	NICKEL  = Coin{"NICKEL", 5}
	DIME    = Coin{"DIME", 10}
	QUARTER = Coin{"QUARTER", 25}
)
