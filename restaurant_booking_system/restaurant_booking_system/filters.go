package restaurant_booking_system

type Cuisine int

const (
	THAI Cuisine = iota
	CHINESE
	NORTHINDIAN
	SOUTHINDIAN
)

var CuisineStrings = map[Cuisine]string{
	THAI:        "THAI",
	CHINESE:     "CHINESE",
	NORTHINDIAN: "NORTHINDIAN",
	SOUTHINDIAN: "SOUTHINDIAN",
}

type Cost int

const (
	GreaterThan10000 Cost = iota
	LessThan10000
)

var CostStrings = map[Cost]string{
	GreaterThan10000: "GreaterThan10000",
	LessThan10000:    "LessThan10000",
}

type CostForTwo int

const (
	Thousand CostForTwo = iota
	FiveHundered
	ThreeThousand
)

var CostForTwoStrings = map[CostForTwo]string{
	Thousand:      "Thousand",
	FiveHundered:  "FiveHundered",
	ThreeThousand: "ThreeThousand",
}

type RestaurantType int

const (
	Veg RestaurantType = iota
	NonVeg
	OpenGarden
	Cafe
	// Add other reservation statuses here if needed
)

var RestaurantTypeStrings = map[RestaurantType]string{
	Veg:        "Veg",
	NonVeg:     "NonVeg",
	OpenGarden: "OpenGarden",
	Cafe:       "Cafe",
}

type Filters struct {
	Location   *Location
	Cuisine    Cuisine
	Cost       Cost
	Type       RestaurantType
	CostForTwo CostForTwo
}
