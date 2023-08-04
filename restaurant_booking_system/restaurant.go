package main

type Restaurant struct {
	RestaurantId   int
	RestaurantName string
	Tables         []*Table
	Location       *Location
	Cost           Cost
	CostForTwo     CostForTwo
	Cuisine        []Cuisine
	Type           RestaurantType
}
