package main

type Location struct {
	City    string
	PinCode int
	Area    string
}

func NewLocation(city string, pinCode int, area string) *Location {
	return &Location{
		City:    city,
		PinCode: pinCode,
		Area:    area,
	}
}
