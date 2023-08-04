package main

import "fmt"

func main() {
	//****************************REGISTER BY OWNER*********************************************
	location := &Location{
		City:    "Patti",
		PinCode: 143416,
		Area:    "Punjab",
	}
	slot1 := &Slot{
		fromTime: 14,
		toTime:   15,
	}
	slot2 := &Slot{
		fromTime: 14,
		toTime:   15,
	}
	tables := []*Table{
		&Table{
			TableId:    1,
			PersonList: nil,
			TableType:  Family,
			Slot:       slot1,
		},
		&Table{
			TableId:    2,
			PersonList: nil,
			TableType:  TwoSeater,
			Slot:       slot2,
		},
	}
	restaurant := &Restaurant{
		RestaurantId:   1,
		RestaurantName: "AANCH",
		Tables:         tables,
		Location:       location,
		Cost:           LessThan10000,
		CostForTwo:     ThreeThousand,
		Cuisine: []Cuisine{
			Cuisine(NorthIndian),
			Cuisine(SouthIndian),
			Cuisine(Chinese),
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
	inventoryManager.AddRestaurant(restaurant)

	//****************************SEARCH BY USER*********************************************
	restaurantList := inventoryManager.getRestaurantList()
	filterBy := &Filters{
		Location: location,
		Cuisine:  NorthIndian,
		Cost:     LessThan10000,
		Type:     Veg,
	}
	searchObj := &Search{
		RestaurantList: restaurantList,
		Filters:        filterBy,
		SearchStrategy: &SearchAlgoStrategy1{},
	}
	filteredRestaurantList := searchObj.SearchRestaurant()
	for _, filteredRestaurant := range filteredRestaurantList {
		fmt.Println(filteredRestaurant)
	}

	//book by user
}
