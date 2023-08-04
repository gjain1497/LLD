package main

type InventoryManager struct {
	RestaurantList []*Restaurant
	OwnerList      []*RestaurantOwner
}

func (im *InventoryManager) RegisterOwner(owner *RestaurantOwner) {
	list := im.OwnerList
	list = append(list, owner)
}

func (im *InventoryManager) AddRestaurant(restaurant *Restaurant) {
	list := im.RestaurantList
	list = append(list, restaurant)
}

func (im *InventoryManager) getRestaurantList() []*Restaurant {
	list := im.RestaurantList
	return list
}
