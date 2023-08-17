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
	err := filteredRestaurantList[0].BookSlot(&booking.Slot{
		Date: "2023-08-22",
		Time: "14:00:00",
	}, 4)
	log.Printf("err1: %v\n", err)
	log.Println()
	if err == nil {
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
	err = filteredRestaurantList[0].BookSlot(&booking.Slot{
		Date: "2023-08-14",
		Time: "12:00:00",
	}, 4)
	log.Printf("err2: %v\n", err)
	log.Println()
	if err == nil {
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
	err = filteredRestaurantList[0].BookSlot(&booking.Slot{
		Date: "2023-08-06",
		Time: "14:00:00",
	}, 4)
	log.Printf("err3: %v\n", err)
	log.Println()
	if err == nil {
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
	filteredRestaurantList[0].AddTimeSlot(
		booking.NewSlot("2023-08-14", "21:00:00"),
	)
	//errChan := make(chan error)
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			err := filteredRestaurantList[0].BookSlot(&booking.Slot{
				Date: "2023-08-14",
				Time: "21:00:00",
			}, 4)
			if err == nil {
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
		Location: &booking.Location{
			City:    "Patti",
			PinCode: 143416,
			Area:    "Punjab",
		},
		Cuisine: booking.NORTHINDIAN,
		Cost:    booking.LessThan10000,
		Type:    booking.Veg,
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
	restaurant, errRes := booking.NewRestaurant(booking.Options{
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
	})
	if errRes != nil {
		fmt.Printf("err initating restaurant: %v\n", errRes)
	}
	var err error
	err = restaurant.AddTimeSlot(
		booking.NewSlot("2023-08-14", "09:00:00"),
	)
	fmt.Printf("err1: %v\n", err)

	err = restaurant.AddTimeSlot(
		booking.NewSlot("2023-08-14", "12:00:00"),
	)
	fmt.Printf("err2: %v\n", err)

	err = restaurant.AddTimeSlot(
		booking.NewSlot("2023-08-14", "18:00:00"),
	)
	fmt.Printf("err3: %v\n", err)
	return restaurant
}
