package splitwise

import (
	"errors"
	"fmt"
	"log"
)

const (
	ErrPercentageNotValid = "sum of percentage should be equal to 100. Please enter valid percentage split"
	ErrSumSharesNotValid  = "sum of individual shares should be equal to total expense. Please enter valid individual exact expenses"
)

type User struct {
	userId      string
	name        string
	mobile      int
	expenseList map[string]float64
}

func init() {
	InitializeUsers() // Initialize users map when the package is imported
}

func NewUser(userId string, username string, mobile int) *User {
	return &User{
		userId:      userId,
		name:        username,
		expenseList: map[string]float64{},
	}
}

var users map[string]*User

func InitializeUsers() {
	users = make(map[string]*User)
}

func AddUser(userId string, userName string, mobile int) {
	user := NewUser(userId, userName, mobile)
	users[userId] = user
}

func AddExpense(amountPaid float64, userId string, numberOfUsers int, userList []string, expenseType string, optionalIntSlice ...[]float64) error {
	userWhoPaid := users[userId]
	switch expenseType {
	case "EQUAL":
		amountRespective := amountPaid / float64(numberOfUsers)
		for _, otherUserID := range userList {
			if otherUserID == userId {
				continue
			}
			userWhoPaid.expenseList[otherUserID] += amountRespective
			users[otherUserID].expenseList[userId] -= amountRespective
		}
	case "EXACT":
		var sumExpense float64
		for _, expense := range optionalIntSlice[0] {
			sumExpense += expense
		}
		if sumExpense < amountPaid {
			return errors.New(ErrSumSharesNotValid)
		}
		for ind, otherUserID := range userList {
			if otherUserID == userId {
				continue
			}
			amountRespective := optionalIntSlice[0][ind]
			userWhoPaid.expenseList[otherUserID] += amountRespective
			users[otherUserID].expenseList[userId] -= amountRespective
		}
	case "PERCENT":
		var sumPercentage float64
		for _, percentage := range optionalIntSlice[0] {
			sumPercentage += percentage
		}
		if int(sumPercentage) < 100 {
			return errors.New(ErrPercentageNotValid)
		}
		for ind, otherUserID := range userList {
			if otherUserID == userId {
				continue
			}
			amountRespective := float64(optionalIntSlice[0][ind]*amountPaid) / float64(100)
			userWhoPaid.expenseList[otherUserID] += amountRespective
			users[otherUserID].expenseList[userId] -= amountRespective
		}
	}
	return nil
}

func ShowBalance(userId string) error {
	var currentUser *User
	if val, ok := users[userId]; ok {
		currentUser = val
	} else {
		return errors.New("map empty")
	}
	mp := currentUser.expenseList

	for user, amount := range mp {
		if amount < 0 {
			continue
		}
		roundedAmount := fmt.Sprintf("%.2f", amount)
		log.Printf("user %s owes user %s : amount %s", user, userId, roundedAmount)
	}
	return nil
}

func ShowBalanceAll() {
	for userId := range users {
		ShowBalance(userId)
	}
}
