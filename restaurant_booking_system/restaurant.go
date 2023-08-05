package main

import "log"

type Restaurant struct {
	RestaurantId   int
	RestaurantName string
	Tables         []*Table
	Location       *Location
	Cost           Cost
	CostForTwo     CostForTwo
	Cuisine        []Cuisine
	Type           RestaurantType
}

func (rest Restaurant) bookTable(table *Table, slot *Slot) {
	for _, tab := range rest.Tables {
		if table.TableId == tab.TableId {
			if table.Booking.Status == AVAILABLE {
				for _, slt := range tab.Slots {
					if slt.FromTime == slot.FromTime && slt.ToTime == slot.ToTime {
						if slt.Status == SLOTAVAILABLE {
							table.Booking.BookingId = 1
							table.Booking.Status = BOOKED
						} else {
							log.Println("Slot not available currently")
							log.Println("Slot is  available from ", slt.FromTime, "till ", slt.ToTime)
						}
					}
				}
			} else {
				log.Println("Table already Booked")
			}
		}
	}
}
