package main

import "log"

type Bike struct {
	Vehicle
	HasStorageBox bool
}

func (b *Bike) Drive() {
	log.Println("Repairing a Bike")
}
