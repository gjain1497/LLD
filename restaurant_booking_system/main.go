package main

import (
	"fmt"
	"log"
)

func main() {
	//****************************REGISTER BY OWNER*********************************************
	location := &Location{
		City:    "Patti",
		PinCode: 143416,
		Area:    "Punjab",
	}
	slottable1List := []*Slot{
		&Slot{
			FromTime: 14,
			ToTime:   15,
			Status:   SLOTAVAILABLE,
		},
		&Slot{
			FromTime: 16,
			ToTime:   17,
			Status:   SLOTAVAILABLE,
		},
	}
	slottable2List := []*Slot{
		&Slot{
			FromTime: 14,
			ToTime:   15,
			Status:   SLOTAVAILABLE,
		},
		&Slot{
			FromTime: 17,
			ToTime:   18,
			Status:   SLOTAVAILABLE,
		},
	}
	tables := []*Table{
		&Table{
			TableId:    1,
			PersonList: nil,
			TableType:  Family,
			Slots:      slottable1List,
			Booking:    &Booking{0, AVAILABLE},
		},
		&Table{
			TableId:    2,
			PersonList: nil,
			TableType:  TwoSeater,
			Slots:      slottable2List,
			Booking:    &Booking{0, AVAILABLE},
		},
	}
	restaurant := &Restaurant{
		RestaurantId:   1,
		RestaurantName: "Aanch",
		Tables:         tables,
		Location:       location,
		Cost:           LessThan10000,
		CostForTwo:     ThreeThousand,
		Cuisine: []Cuisine{
			Cuisine(NORTHINDIAN),
			Cuisine(SOUTHINDIAN),
			Cuisine(CHINESE),
		},
		Type: Veg,
	}
	owner := &RestaurantOwner{
		OwnerId:    123,
		OwnerName:  "Ankush",
		Restaurant: restaurant,
	}
	inventoryManager := InventoryManager{}
	inventoryManager.RegisterOwner(owner) //it will add owner to the list of owners
	inventoryManager.AddRestaurant(restaurant, owner.OwnerId)

	//****************************SEARCH BY USER*********************************************
	restaurantList := inventoryManager.getRestaurantList()
	displayRestaurants(restaurantList)

	filterBy := &Filters{
		Location: location,
		Cuisine:  NORTHINDIAN,
		Cost:     LessThan10000,
		Type:     Veg,
	}

	searchObj := &Search{
		//RestaurantList: restaurantList,
		Filters: filterBy,
		//SearchStrategy: &SearchAlgoStrategy1{},
	}
	filteredRestaurantList := searchObj.SearchRestaurant(Filters)
	log.Println()
	log.Println("RESTAURANTS AFTER FILTERING")
	log.Println()
	displayRestaurants(filteredRestaurantList)
	log.Println()

	////****************************BOOK TABLE BY USER*********************************************

	//restaurant -> tables -> slots

	selectedTable := tables[1]        //table2
	selectedSlot := slottable2List[1] //slottable2 -> 2sra slot
	filteredRestaurantList[0].bookTable(date, selectedSlot, numberofBande)
	log.Println()
	log.Println("STATUS AFTER BOOKING")
	log.Println()
	displayRestaurants(restaurantList)
	log.Println()

	////****************************BOOK AGAIN SAME TABLE BY USER*********************************************
	selectedTableAgain := tables[1]
	selectedSlotAgain := slottable2List[1]
	filteredRestaurantList[0].bookTable(selectedTableAgain, selectedSlotAgain)
	log.Println()
	log.Println("STATUS AFTER BOOKING")
	log.Println()
	displayRestaurants(restaurantList)
	log.Println()
}

func displayRestaurants(restaurantList []*Restaurant) {
	fmt.Printf("restaurantList %+v", restaurantList)
	//for _, rest := range restaurantList {
	//	log.Println("RestaurantName ", rest.RestaurantName)
	//	for _, table := range rest.Tables {
	//		log.Println("table_id: ", table.TableId)
	//		log.Println("table_type: ", TableTypeStrings[table.TableType])
	//		log.Println("table_booking_status: ", bookingStatusStrings[table.Booking.Status])
	//		for _, slot := range table.Slots {
	//			log.Println("slot_status: ", SlotStatusString[slot.Status], "from_: ", slot.FromTime, "to_: ", slot.ToTime)
	//		}
	//	}
	//	log.Println("City: ", rest.Location.City, ",Area: ", rest.Location.Area, ",PinCode: ", rest.Location.PinCode)
	//	log.Println("Cost_Range: ", CostStrings[rest.Cost])
	//	log.Println("Cost For Two: ", CostForTwoStrings[rest.CostForTwo])
	//	for _, cus := range rest.Cuisine {
	//		log.Print("Cuisine: ", CuisineStrings[cus]+" ")
	//	}
	//	log.Println("Type: ", RestaurantTypeStrings[rest.Type])
	//}
}
