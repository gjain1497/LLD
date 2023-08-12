package main

import "log"

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

func AddRestaurant(restaurant *Restaurant) {
	restaurants = append(restaurants, restaurant)
}

func GetAllRestaurants() []*Restaurant {
	return restaurants
}

func InitRestaurant() *Restaurant {
	// Initialize the list of restaurants from your data source (e.g., database, API, etc.)
	location := &Location{
		City:    "Patti",
		PinCode: 143416,
		Area:    "Punjab",
	}
	slotList := []*Slot{
		&Slot{
			Date:           12,
			Time:           1,
			Status:         SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
		&Slot{
			Date:           12,
			Time:           2,
			Status:         SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
		&Slot{
			Date:           12,
			Time:           3,
			Status:         SLOTAVAILABLE,
			NumberOfPeople: 4,
		},
	}
	restaurant := &Restaurant{
		RestaurantId:   1,
		RestaurantName: "Aanch",
		NumberOfTables: 4,
		Slots:          slotList,
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
	return restaurant
}

func (r *Restaurant) AddAvailableSlot(slot *Slot) {
	r.Slots = append(r.Slots, slot)
}

func (r *Restaurant) bookTable(slot *Slot, people int) bool {
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
