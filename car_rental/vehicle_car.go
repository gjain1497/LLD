package main

import "log"

type Car struct {
	Vehicle
	NumberOfDoors int
	IsConvertible bool
}

func (c *Car) Drive() {
	log.Println("Repairing a Car")
}
