package restaurant_booking_system

import (
	"errors"
	"log"
	"sync"
	"time"
)

const (
	maxDiffAdvanceBooking = time.Hour * 24 * 7
	// if I change the year to any other value: 2001-5/7, it isn't working. why?
	dateParser = "2006-01-02"
	timeParser = "2006-01-02 15:04:05"
)
const (
	ErrBookingOutOfRange       = "booking only allowed for upto 7 days in advance and only for future time"
	ErrHoursOutOfRange         = "hours should be within range of 0 to 23"
	ErrInsufficientTableInSlot = "insufficient tables in the current time slot for booking"
)

type Restaurant struct {
	RestaurantId   int
	RestaurantName string
	Slots          []*Slot
	Location       *Location
	Cost           Cost
	CostForTwo     CostForTwo
	Cuisine        []Cuisine
	Type           RestaurantType
	mutex          sync.Mutex
}

var restaurants []*Restaurant

func AddRestaurant(r *Restaurant) {
	restaurants = append(restaurants, r)
}

func GetAllRestaurants() []*Restaurant {
	return restaurants
}

//func DisplayRestaurants(restaurantList []*Restaurant) {
//	fmt.Printf("\n Restaurant List Result: ")
//	for _, r := range restaurantList {
//		fmt.Printf("\n%+v", *r)
//	}
//	fmt.Println()
//}

func DisplayRestaurants(restaurantList []*Restaurant) {
	for _, rest := range restaurantList {
		log.Println("RestaurantName ", rest.RestaurantName)
		for i, slot := range rest.Slots {
			log.Println("number of tables in slot:", i, "=", slot.NumberOfTables, "slot_date: ", slot.Date, "slot_time: ", slot.Time)
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

func (r *Restaurant) AddTimeSlot(slot *Slot) error {
	//if slot.Time < 0 || slot.Time > 23 {
	//	return errors.New(ErrHoursOutOfRange)
	//}
	r.Slots = append(r.Slots, slot)
	return nil
}

func SearchRestaurant(f *Filters) []*Restaurant {
	restSearchObj := &RestaurantSearch{
		RestaurantList:           restaurants,
		RestaurantSearchStrategy: &SearchAlgoStrategyElasticIndexing1{},
	}
	filteredRestList := restSearchObj.SearchRestaurant(f)
	return filteredRestList
}

func (r *Restaurant) BookSlot(slot *Slot, people int) (bool, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	DateTime := slot.Date + " " + slot.Time
	parsedBookingTime, err := time.Parse(timeParser, DateTime)
	if err != nil {
		return false, err
	}
	//log.Println("parsedBookingTime ", parsedBookingTime)
	//log.Println("time Until parsedBookingTime ", time.Until(parsedBookingTime))
	//log.Println("maxDiffAdvanceBooking ", maxDiffAdvanceBooking)
	if time.Until(parsedBookingTime) > maxDiffAdvanceBooking || time.Until(parsedBookingTime) < 0 {
		return false, errors.New(ErrBookingOutOfRange)
	}
	if slot.NumberOfTables != 0 { //assuming (people int) are adjusted in 1 table, so 1 table gone
		slot.NumberOfTables--
		return true, nil
	} else {
		return false, errors.New(ErrInsufficientTableInSlot)
	}
}
