package main

type Cuisine int

const (
	Thai Cuisine = iota
	Chinese
	NorthIndian
	SouthIndian
)

type Cost int

const (
	GreaterThan10000 Cost = iota
	LessThan10000
)

type CostForTwo int

const (
	Thousand CostForTwo = iota
	FiveHundered
	ThreeThousand
)

type RestaurantType int

const (
	Veg RestaurantType = iota
	NonVeg
	OpenGarden
	Cafe
	// Add other reservation statuses here if needed
)

type Filters struct {
	Location   *Location
	Cuisine    Cuisine
	Cost       Cost
	Type       RestaurantType
	CostForTwo CostForTwo
}

func (f Filters) GetRestaurantList() {
	//var filteredRestaurantList []*Restaurant
	//restaurantList := f.RestaurantList
	//for _,i := range restaurantList{
	//	if restaurantList[i].Location
	//}
}
