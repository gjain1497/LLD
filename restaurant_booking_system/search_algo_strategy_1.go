package main

type SearchAlgoStrategy1 struct {
}

func (s1 *SearchAlgoStrategy1) searchRestaurant(restaurantList []*Restaurant, filterBy *Filters) []*Restaurant {
	//use strategy1 to filter restaurant
	var filteredRestaurantList []*Restaurant
	for _, restaurant := range restaurantList {
		if filterBy.Location == restaurant.Location &&
			filterBy.Cost == restaurant.Cost &&
			filterBy.CostForTwo == restaurant.CostForTwo {
			filteredRestaurantList = append(filteredRestaurantList, restaurant)
		}
	}
	return filteredRestaurantList
}
