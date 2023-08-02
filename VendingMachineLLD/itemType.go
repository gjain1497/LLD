package main

// Define a custom struct to represent the Coin enum-like behavior.
type ItemType struct {
	name string
}

// Declare constants representing each coin.
var (
	COKE  = ItemType{"COKE"}
	PEPSI = ItemType{"PEPSI"}
	JUICE = ItemType{"JUICE"}
	SODA  = ItemType{"SODA"}
)
