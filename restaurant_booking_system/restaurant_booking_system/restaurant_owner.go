package restaurant_booking_system

type RestaurantOwner struct {
	OwnerId     int
	OwnerName   string
	Restaurants []*Restaurant
}

func (r *RestaurantOwner) RegisterRestaurant(restaurant *Restaurant) {
	r.Restaurants = append(r.Restaurants, restaurant)
	AddRestaurant(restaurant) //as soon as owner has added a restaurant add that in restaurant system also
}
