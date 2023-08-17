package restaurant_booking_system

import (
	"errors"
	"fmt"
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
	ErrInsufficientTableInSlot = "insufficient tables in the current time slot for booking"
)

type Options struct {
	RestaurantId   int
	RestaurantName string
	Location       *Location
	Cost           Cost
	CostForTwo     CostForTwo
	Cuisine        []Cuisine
	Type           RestaurantType
} //options bs filter krna tha isliye use kra ya vaise bhi private vale concept ke lie bhi

type Restaurant struct {
	restaurantId   int
	restaurantName string
	slots          []*Slot
	location       *Location
	cost           Cost
	costForTwo     CostForTwo
	cuisine        []Cuisine
	resType        RestaurantType
	mutex          sync.Mutex
}

func NewRestaurant(o Options) (*Restaurant, error) {

	r := &Restaurant{
		restaurantId:   o.RestaurantId,
		restaurantName: o.RestaurantName,
		cost:           o.Cost,
		costForTwo:     o.CostForTwo,
		cuisine:        o.Cuisine,
		resType:        o.Type,
	}
	return r, nil
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
		log.Println("RestaurantName ", rest.restaurantName)
		for i, slot := range rest.slots {
			log.Println("number of tables in slot:", i, "=", slot.NumberOfTables, "slot_date: ", slot.Date, "slot_time: ", slot.Time)
		}
		log.Println("City: ", rest.location.City, ",Area: ", rest.location.Area, ",PinCode: ", rest.location.PinCode)
		log.Println("Cost_Range: ", CostStrings[rest.cost])
		log.Println("Cost For Two: ", CostForTwoStrings[rest.costForTwo])
		for _, cus := range rest.cuisine {
			log.Print("Cuisine: ", CuisineStrings[cus]+" ")
		}
		log.Println("Type: ", RestaurantTypeStrings[rest.resType])
	}
}

func (r *Restaurant) AddTimeSlot(slot *Slot) error {
	//if slot.Time < 0 || slot.Time > 23 {
	//	return errors.New(ErrHoursOutOfRange)
	//}
	r.slots = append(r.slots, slot)
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

func (r *Restaurant) BookSlot(slot *Slot, people int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	DateTime := fmt.Sprintf("%s %s", slot.Date, slot.Time)
	parsedBookingTime, err := time.Parse(timeParser, DateTime)
	if err != nil {
		return err
	}
	//log.Println("parsedBookingTime ", parsedBookingTime)
	//log.Println("time Until parsedBookingTime ", time.Until(parsedBookingTime))
	//log.Println("maxDiffAdvanceBooking ", maxDiffAdvanceBooking)
	if time.Until(parsedBookingTime) > maxDiffAdvanceBooking || time.Until(parsedBookingTime) < 0 {
		return errors.New(ErrBookingOutOfRange)
	}
	if slot.NumberOfTables != 0 { //assuming (people int) are adjusted in 1 table, so 1 table gone
		slot.NumberOfTables--
		return nil
	}
	return errors.New(ErrInsufficientTableInSlot)
}
