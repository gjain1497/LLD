package main

import (
	booking "github.com/gjain1497/LLD/restaurant_booking_system/restaurant_booking_system"
	"log"
)

func main() {
	//****************************INITIALISE RESTAURANT*********************************************
	rest := InitRestaurant()

	//****************************REGISTER BY OWNER*********************************************
	owner := &booking.RestaurantOwner{
		OwnerId:   123,
		OwnerName: "Ankush Garg",
	}
	owner.RegisterRestaurant(rest)
	restaurantList := booking.GetAllRestaurants()

	//****************************DISPLAY ALL RESTAURANTS PRESENT IN THE SYSTEM*********************************************
	booking.DisplayRestaurants(restaurantList)

	//****************************SEARCH BY USER BY APPLYING VARIOUS FILTERS *********************************************
	filterBy := &booking.Filters{
		Location: rest.Location,
		Cuisine:  booking.NORTHINDIAN,
		Cost:     booking.LessThan10000,
		Type:     booking.Veg,
	}
	//
	//restSearchObj := &booking.RestaurantSearch{ //this is not the correct way  as RestaurantSearchStrategy
	//will not be known to user, it should be somehow handled internally in the code and not exposed in the main function
	//	RestaurantList:           restaurantList,
	//	RestaurantSearchStrategy: &SearchAlgoStrategy1{},
	//}
	filteredRestaurantList := booking.SearchRestaurant(filterBy)
	log.Println()
	log.Println("RESTAURANTS AFTER FILTERING")
	log.Println()
	booking.DisplayRestaurants(restaurantList)
	log.Println()

	////****************************BOOK TABLE BY USER*********************************************
	selectedSlot := rest.Slots[0]
	log.Println("TRYING TO  BOOK A SLOT")
	log.Println()
	booked := filteredRestaurantList[0].BookTable(selectedSlot, 4)
	log.Println()
	if booked {
		log.Println("STATUS AFTER BOOKING")
		log.Println()
		booking.DisplayRestaurants(restaurantList)
		log.Println()
	}

	////****************************BOOK AGAIN SAME TABLE BY USER*********************************************
	selectedSlotAgain := rest.Slots[0]
	log.Println("TRYING TO  BOOK A SLOT")
	log.Println()
	bookedAgain := filteredRestaurantList[0].BookTable(selectedSlotAgain, 4)
	log.Println()
	if bookedAgain {
		log.Println("STATUS AFTER BOOKING")
		log.Println()
		booking.DisplayRestaurants(restaurantList)
		log.Println()
	}
}

func InitRestaurant() *booking.Restaurant {
	// Initialize the list of restaurants from your data source (e.g., database, API, etc.)
	location := &booking.Location{
		City:    "Patti",
		PinCode: 143416,
		Area:    "Punjab",
	}
	slotList := []*booking.Slot{
		{
			Date:           12,
			Time:           1,
			Status:         booking.SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
		{
			Date:           12,
			Time:           2,
			Status:         booking.SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
		{
			Date:           12,
			Time:           3,
			Status:         booking.SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
	}
	restaurant := &booking.Restaurant{
		RestaurantId:   1,
		RestaurantName: "Aanch",
		NumberOfTables: 4,
		Slots:          slotList,
		Location:       location,
		Cost:           booking.LessThan10000,
		CostForTwo:     booking.ThreeThousand,
		Cuisine: []booking.Cuisine{
			booking.NORTHINDIAN,
			booking.SOUTHINDIAN,
			booking.CHINESE,
		},
		Type: booking.Veg,
	}
	return restaurant
}
