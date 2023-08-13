package main

import (
	"fmt"
	booking "github.com/gjain1497/LLD/restaurant_booking_system/restaurant_booking_system"
	"log"
	"sync"
	"time"
)

func bookTablesSyncHelper(filteredRestaurantList []*booking.Restaurant) {

	//1st case
	log.Println("OUT OF RANGE TIME FUTURE")
	log.Println()
	booked1, err := filteredRestaurantList[0].BookSlot(&booking.Slot{
		Date: "2023-08-20",
		Time: "14:00:00",
	}, 4)
	log.Printf("err1: %v\n", err)
	log.Println()
	if booked1 {
		log.Println("SEATS BOOKED SUCCESSFULLY")
	} else {
		log.Println("SEATS NOT BOOKED")
	}
	log.Println()
	log.Println("STATUS AFTER BOOKING")
	log.Println()
	booking.DisplayRestaurants(filteredRestaurantList)
	log.Println()

	//2nd case
	log.Println("IN RANGE TIME")
	log.Println()
	booked2, err := filteredRestaurantList[0].BookSlot(filteredRestaurantList[0].Slots[1], 4)
	log.Printf("err2: %v\n", err)
	log.Println()
	if booked2 {
		log.Println("SEATS BOOKED SUCCESSFULLY")
	} else {
		log.Println("SEATS NOT BOOKED")
	}
	log.Println()
	log.Println("STATUS AFTER BOOKING")
	log.Println()
	booking.DisplayRestaurants(filteredRestaurantList)
	log.Println()

	//3rd case
	log.Println("OUT OF RANGE TIME PAST")
	log.Println()
	booked3, err := filteredRestaurantList[0].BookSlot(&booking.Slot{
		Date: "2023-08-06",
		Time: "14:00:00",
	}, 4)
	log.Printf("err3: %v\n", err)
	log.Println()
	if booked3 {
		log.Println("SEATS BOOKED SUCCESSFULLY")
	} else {
		log.Println("SEATS NOT BOOKED")
	}
	log.Println()
	log.Println("STATUS AFTER BOOKING")
	log.Println()
	booking.DisplayRestaurants(filteredRestaurantList)
	log.Println()
}

func bookTablesConcurrentHelper(filteredRestaurantList []*booking.Restaurant) {
	filteredRestaurantList[0].AddTimeSlot(&booking.Slot{
		Date:           "2023-08-13",
		Time:           "19:00:00",
		NumberOfTables: 10,
	})
	//errChan := make(chan error)
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			isBooked, _ := filteredRestaurantList[0].BookSlot(filteredRestaurantList[0].Slots[1], 4)
			if isBooked {
				log.Println("Slot booked for candidate ", num)
			} else {
				log.Println("Slot not booked for candidate ", num)
			}
			//errChan <- err
		}(i)
	}
	wg.Wait()
	time.Sleep(5 * time.Second) // Wait for output to display

	//for i := 0; i < 10; i++ {
	//	err := <-errChan
	//	fmt.Println("Err concurrent: ", err)
	//}
}

func main() {
	//****************************INITIALISE RESTAURANT*********************************************
	rest := InitRestaurant()

	//****************************REGISTER BY OWNER*********************************************
	owner := &booking.RestaurantOwner{
		OwnerId:   123,
		OwnerName: "Ankush Garg",
	}
	owner.RegisterRestaurant(rest)

	//****************************DISPLAY ALL RESTAURANTS PRESENT IN THE SYSTEM*********************************************
	restaurantList := booking.GetAllRestaurants()
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

	//****************************BOOK TABLES SEQUENTIALLY *********************************************
	bookTablesSyncHelper(filteredRestaurantList)
	//booking.DisplayRestaurants(filteredRestaurantList)

	//****************************BOOK TABLES CONCURRENTLY *********************************************
	bookTablesConcurrentHelper(filteredRestaurantList)
	booking.DisplayRestaurants(filteredRestaurantList)
}

func InitRestaurant() *booking.Restaurant {
	// Initialize the list of restaurants from your data source (e.g., database, API, etc.)
	location := &booking.Location{
		City:    "Patti",
		PinCode: 143416,
		Area:    "Punjab",
	}
	restaurant := &booking.Restaurant{
		RestaurantId:   1,
		RestaurantName: "Aanch",
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
	var err error
	err = restaurant.AddTimeSlot(&booking.Slot{
		Date:           "2023-08-13",
		Time:           "09:00:00",
		NumberOfTables: 10,
	})
	fmt.Printf("err1: %v\n", err)

	err = restaurant.AddTimeSlot(&booking.Slot{
		Date:           "2023-08-13",
		Time:           "12:00:00",
		NumberOfTables: 10,
	})
	fmt.Printf("err2: %v\n", err)

	err = restaurant.AddTimeSlot(&booking.Slot{
		Date:           "2023-08-13",
		Time:           "18:00:00",
		NumberOfTables: 10,
	})
	fmt.Printf("err3: %v\n", err)
	return restaurant
}
