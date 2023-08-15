package restaurant_booking_system

import "log"

type SearchAlgoStrategyElasticIndexing1 struct {
}

func (s1 *SearchAlgoStrategyElasticIndexing1) searchRestaurant(restaurantList []*Restaurant, filterBy *Filters) []*Restaurant {
	//use strategy1 to filter restaurant
	log.Println("Reached in SearchAlgoStrategyElasticIndexing1")
	var filteredRestaurantList []*Restaurant
	for _, restaurant := range restaurantList {
		if (filterBy.Location.City == "" || filterBy.Location.City == restaurant.location.City) &&
			(filterBy.Cost == 0 || filterBy.Cost == restaurant.cost) &&
			(filterBy.Type == 0 || filterBy.Type == restaurant.resType) {
			filteredRestaurantList = append(filteredRestaurantList, restaurant)
		}
	}
	return filteredRestaurantList
}
