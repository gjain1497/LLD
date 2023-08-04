package main

type VehicleRentalSystem struct {
	storeList []*Store
	userList  []*User
}

func NewVehicleRentalSystem(stores []*Store, users []*User) *VehicleRentalSystem {
	return &VehicleRentalSystem{
		storeList: stores,
		userList:  users,
	}
}

func (vrs *VehicleRentalSystem) getStore(location *Location) *Store {
	// Implement logic to find and filter out the store based on location
	// For now, return the first store in the list
	storeList := vrs.StoreList()
	return storeList[0]
}

func (vrs *VehicleRentalSystem) StoreList() []*Store {
	return vrs.storeList
}

func (vrs *VehicleRentalSystem) SetStoreList(storeList []*Store) {
	vrs.storeList = storeList
}

func (vrs *VehicleRentalSystem) UserList() []*User {
	return vrs.userList
}

func (vrs *VehicleRentalSystem) SetUserList(userList []*User) {
	vrs.userList = userList
}

// Additional methods to add and remove users and stores can be added here
