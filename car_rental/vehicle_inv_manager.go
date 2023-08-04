package main

type VehicleInventoryManagement struct {
	vehicles []*Vehicle
}

func NewVehicleInventoryManagement(vehicles []*Vehicle) *VehicleInventoryManagement {
	return &VehicleInventoryManagement{
		vehicles: vehicles,
	}
}

func (vim *VehicleInventoryManagement) GetVehicles(vehicleType VehicleType) []*Vehicle {
	// Filtering logic if required
	return vim.vehicles
}

func (vim *VehicleInventoryManagement) SetVehicles(vehicles []*Vehicle) {
	vim.vehicles = vehicles
}
