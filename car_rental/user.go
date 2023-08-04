package main

type User struct {
	UserID         int
	UserName       int
	DrivingLicense int
}

// Getters and Setters (using Go's convention)

func (u *User) GetUserID() int {
	return u.UserID
}

func (u *User) SetUserID(userID int) {
	u.UserID = userID
}

func (u *User) GetUserName() int {
	return u.UserName
}

func (u *User) SetUserName(userName int) {
	u.UserName = userName
}

func (u *User) GetDrivingLicense() int {
	return u.DrivingLicense
}

func (u *User) SetDrivingLicense(drivingLicense int) {
	u.DrivingLicense = drivingLicense
}
