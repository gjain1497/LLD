package restaurant_booking_system

type SearchAlgoStrategy interface {
	searchRestaurant(restaurantList []*Restaurant, filterBy *Filters) []*Restaurant
}
