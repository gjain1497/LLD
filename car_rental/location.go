package main

type Location struct {
	address string
	pincode int
	city    string
	state   string
	country string
}

func NewLocation(pincode int, city, state, country string) *Location {
	location := &Location{
		pincode: pincode,
		city:    city,
		state:   state,
	}
	return location
}
