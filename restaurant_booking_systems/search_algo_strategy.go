package main

type SearchAlgoStrategy interface {
	searchRestaurant(restaurantList []*Restaurant, filterBy *Filters) []*Restaurant
}
