package main

import (
	"log"
	"time"
)

type VehicleType int

const (
	Caar VehicleType = iota
	Truck
	Van
	Motorcycle
)

type Status int

const (
	Available Status = iota
	Booked
	NotAvailable
)

type Vehicle struct {
	VehicleID         int
	VehicleNumber     int
	VehicleType       VehicleType
	CompanyName       string
	ModelName         string
	KmDriven          int
	ManufacturingDate time.Time
	Average           int
	CC                int
	DailyRentalCost   int
	HourlyRentalCost  int
	NoOfSeat          int
	Status            Status
}

// Getters and Setters (using Go's convention)

func (v *Vehicle) GetVehicleID() int {
	return v.VehicleID
}

func (v *Vehicle) SetVehicleID(vehicleID int) {
	v.VehicleID = vehicleID
}

func (v *Vehicle) GetVehicleNumber() int {
	return v.VehicleNumber
}

func (v *Vehicle) SetVehicleNumber(vehicleNumber int) {
	v.VehicleNumber = vehicleNumber
}

func (v *Vehicle) GetVehicleType() VehicleType {
	return v.VehicleType
}

func (v *Vehicle) SetVehicleType(vehicleType VehicleType) {
	v.VehicleType = vehicleType
}

func (v *Vehicle) GetCompanyName() string {
	return v.CompanyName
}

func (v *Vehicle) SetCompanyName(companyName string) {
	v.CompanyName = companyName
}

func (v *Vehicle) GetModelName() string {
	return v.ModelName
}

func (v *Vehicle) SetModelName(modelName string) {
	v.ModelName = modelName
}

func (v *Vehicle) GetKmDriven() int {
	return v.KmDriven
}

func (v *Vehicle) SetKmDriven(kmDriven int) {
	v.KmDriven = kmDriven
}

func (v *Vehicle) GetManufacturingDate() time.Time {
	return v.ManufacturingDate
}

func (v *Vehicle) SetManufacturingDate(manufacturingDate time.Time) {
	v.ManufacturingDate = manufacturingDate
}

func (v *Vehicle) GetAverage() int {
	return v.Average
}

func (v *Vehicle) SetAverage(average int) {
	v.Average = average
}

func (v *Vehicle) GetCC() int {
	return v.CC
}

func (v *Vehicle) SetCC(cc int) {
	v.CC = cc
}

func (v *Vehicle) GetDailyRentalCost() int {
	return v.DailyRentalCost
}

func (v *Vehicle) SetDailyRentalCost(dailyRentalCost int) {
	v.DailyRentalCost = dailyRentalCost
}

func (v *Vehicle) GetHourlyRentalCost() int {
	return v.HourlyRentalCost
}

func (v *Vehicle) SetHourlyRentalCost(hourlyRentalCost int) {
	v.HourlyRentalCost = hourlyRentalCost
}

func (v *Vehicle) GetNoOfSeat() int {
	return v.NoOfSeat
}

func (v *Vehicle) SetNoOfSeat(noOfSeat int) {
	v.NoOfSeat = noOfSeat
}

func (v *Vehicle) GetStatus() Status {
	return v.Status
}

func (v *Vehicle) SetStatus(status Status) {
	v.Status = status
}

func (v *Vehicle) Drive() {
	log.Println("Repairing a Vehicle")
}
