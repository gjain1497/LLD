package main

type Search struct {
	RestaurantList []*Restaurant
	Filters        *Filters
	SearchStrategy SearchAlgoStrategy
}

func NewSearch(restaurantList []*Restaurant, filters *Filters, searchStrategy SearchAlgoStrategy) *Search {
	return &Search{
		RestaurantList: restaurantList,
		Filters:        filters,
		SearchStrategy: searchStrategy,
	}
}

func (s *Search) SearchRestaurant(filters Filters) []*Restaurant {
	//now SearchStrategy can be any of SearchStrategy1 or SearchStrategy2 based on whatever we pass in the request
	return s.SearchStrategy.searchRestaurant(s.RestaurantList, s.Filters)
}
