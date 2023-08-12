package main

type RestaurantOwner struct {
	OwnerId    int
	OwnerName  string
	Restaurant *Restaurant
}

func NewRestaurantOwner(ownerId int, ownerName string, restaurant *Restaurant) *RestaurantOwner {
	return &RestaurantOwner{
		OwnerId:    ownerId,
		OwnerName:  ownerName,
		Restaurant: restaurant,
	}
}
