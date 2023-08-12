package main

type RestaurantOwner struct {
	OwnerId     int
	OwnerName   string
	Restaurants []*Restaurant
}

func (r *RestaurantOwner) RegisterRestaurant(restaurant *Restaurant) {
	r.Restaurants = append(r.Restaurants, restaurant)
}
