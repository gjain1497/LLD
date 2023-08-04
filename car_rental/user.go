package main

type User struct {
	UserID         int
	UserName       string
	DrivingLicense int
}

// Getters and Setters (using Go's convention)

func (u *User) GetUserID() int {
	return u.UserID
}

func (u *User) SetUserID(userID int) {
	u.UserID = userID
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}

func (u *User) GetDrivingLicense() int {
	return u.DrivingLicense
}

func (u *User) SetDrivingLicense(drivingLicense int) {
	u.DrivingLicense = drivingLicense
}
