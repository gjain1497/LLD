package main

type Location struct {
	Pincode int
	City    string
	State   string
	Country string
}

func NewLocation(pincode int, city, state, country string) *Location {
	location := &Location{
		Pincode: pincode,
		City:    city,
		State:   state,
		Country: country,
	}
	return location
}
