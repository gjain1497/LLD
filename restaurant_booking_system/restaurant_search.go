package main

type RestaurantSearch struct {
	RestaurantList           []*Restaurant
	RestaurantSearchStrategy SearchAlgoStrategy
}

func (s *RestaurantSearch) SearchRestaurant(filters *Filters) []*Restaurant {
	//now SearchStrategy can be any of SearchStrategy1 or SearchStrategy2 based on whatever we pass in the request
	return s.RestaurantSearchStrategy.searchRestaurant(s.RestaurantList, filters)
}
