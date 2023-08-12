package restaurant_booking_system

import (
	"log"
)

type Restaurant struct {
	RestaurantId   int
	RestaurantName string
	NumberOfTables int
	Slots          []*Slot
	Location       *Location
	Cost           Cost
	CostForTwo     CostForTwo
	Cuisine        []Cuisine
	Type           RestaurantType
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

func (r *Restaurant) AddAvailableSlot(slot *Slot) {
	r.Slots = append(r.Slots, slot)
}

func SearchRestaurant(f *Filters) []*Restaurant {
	restSearchObj := &RestaurantSearch{
		RestaurantList:           restaurants,
		RestaurantSearchStrategy: &SearchAlgoStrategyElasticIndexing1{},
	}
	filteredRestList := restSearchObj.SearchRestaurant(f)
	return filteredRestList
}

func (r *Restaurant) BookTable(slot *Slot, people int) bool {
	if slot.Status == SLOTAVAILABLE {
		slot.Status = SLOTNOTAVAILABLE
		slot.NumberOfPeople = people
		r.NumberOfTables--
		return true
	} else {
		log.Println("Slot not available currently")
		return false
	}
}
