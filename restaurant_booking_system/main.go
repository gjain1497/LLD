package main

import (
	"fmt"
	"log"
)

func main() {
	//****************************REGISTER BY OWNER*********************************************
	rest := InitRestaurant()
	AddRestaurant(rest)
	restaurantList := GetAllRestaurants()
	fmt.Println("len restaurantList ", len(restaurantList))

	owner := &RestaurantOwner{
		OwnerId:   123,
		OwnerName: "Ankush Garg",
	}
	owner.RegisterRestaurant(restaurantList[0])

	//****************************SEARCH BY USER*********************************************
	displayRestaurants(restaurantList)

	filterBy := &Filters{
		Location: rest.Location,
		Cuisine:  NORTHINDIAN,
		Cost:     LessThan10000,
		Type:     Veg,
	}

	restSearchObj := &RestaurantSearch{
		RestaurantList:           restaurantList,
		RestaurantSearchStrategy: &SearchAlgoStrategy1{},
	}

	filteredRestaurantList := restSearchObj.SearchRestaurant(filterBy)
	log.Println()
	log.Println("RESTAURANTS AFTER FILTERING")
	log.Println()
	displayRestaurants(filteredRestaurantList)
	log.Println()

	////****************************BOOK TABLE BY USER*********************************************

	selectedSlot := rest.Slots[0]
	log.Println("TRYING TO  BOOK A SLOT")
	log.Println()
	booked := filteredRestaurantList[0].bookTable(selectedSlot, 4)
	log.Println()
	if booked {
		log.Println("STATUS AFTER BOOKING")
		log.Println()
		displayRestaurants(restaurantList)
		log.Println()
	}

	////****************************BOOK AGAIN SAME TABLE BY USER*********************************************
	selectedSlotAgain := rest.Slots[0]
	log.Println("TRYING TO  BOOK A SLOT")
	log.Println()
	bookedAgain := filteredRestaurantList[0].bookTable(selectedSlotAgain, 4)
	log.Println()
	if bookedAgain {
		log.Println("STATUS AFTER BOOKING")
		log.Println()
		displayRestaurants(restaurantList)
		log.Println()
	}
}

func displayRestaurants(restaurantList []*Restaurant) {
	for _, rest := range restaurantList {
		log.Println("RestaurantName ", rest.RestaurantName)
		log.Println("Number of tables ", rest.NumberOfTables)
		for _, slot := range rest.Slots {
			log.Println("slot_status: ", SlotStatusString[slot.Status], "slot_date: ", slot.Date, "slot_time: ", slot.Time, "seats: ", slot.NumberOfPeople)
		}
		log.Println("City: ", rest.Location.City, ",Area: ", rest.Location.Area, ",PinCode: ", rest.Location.PinCode)
		log.Println("Cost_Range: ", CostStrings[rest.Cost])
		log.Println("Cost For Two: ", CostForTwoStrings[rest.CostForTwo])
		for _, cus := range rest.Cuisine {
			log.Print("Cuisine: ", CuisineStrings[cus]+" ")
		}
		log.Println("Type: ", RestaurantTypeStrings[rest.Type])
	}
}
