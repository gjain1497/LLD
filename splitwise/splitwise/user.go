package splitwise

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
)

const (
	ErrPercentageNotValid  = "sum of percentage should be equal to 100. Please enter valid percentage split"
	ErrSumSharesNotValid   = "sum of individual shares should be equal to total expense. Please enter valid individual exact expenses"
	ErrExpenseNameNotFound = "expense name not found"
)

type User struct {
	userId                    string
	name                      string
	mobile                    int
	userPaymentList           map[string]float64
	simplifiedUserPaymentList map[string]float64
}

func init() {
	InitializeUsers() // Initialize users map when the package is imported
}

func NewUser(userId string, username string, mobile int) *User {
	return &User{
		userId:                    userId,
		name:                      username,
		mobile:                    mobile,
		userPaymentList:           map[string]float64{},
		simplifiedUserPaymentList: map[string]float64{},
	}
}

var usersMapping map[string]*User
var allExpenseList []string
var expenseNameMapping map[string]map[string]float64

func InitializeUsers() {
	usersMapping = make(map[string]*User)
	allExpenseList = make([]string, 0)
	expenseNameMapping = make(map[string]map[string]float64)
}

func AddUser(userId string, userName string, mobile int) {
	user := NewUser(userId, userName, mobile)
	usersMapping[userId] = user
}
func adjustExpenseRespective(userWhoPaid *User, userList []string, userId string, expenseName string, amountRespective []float64) error {
	for ind, otherUserID := range userList {
		if amountRespective[ind] == math.SmallestNonzeroFloat64 {
			continue
		}
		userWhoPaid.userPaymentList[otherUserID] += amountRespective[ind]
		usersMapping[otherUserID].userPaymentList[userId] -= amountRespective[ind]
		expenseNameMapping[expenseName][otherUserID] = amountRespective[ind]
	}
	return nil
}
func AddExpense(expenseName string, amountPaid float64, userId string, numberOfUsers int, userList []string, expenseType string, optionalIntSlice ...[]float64) error {
	userWhoPaid := usersMapping[userId]
	allExpenseList = append(allExpenseList, expenseName)
	if _, exists := expenseNameMapping[expenseName]; !exists {
		expenseNameMapping[expenseName] = make(map[string]float64)
	}
	expenseNameMapping[expenseName][userId] = amountPaid

	switch expenseType {
	case "EQUAL":
		var amountRespective []float64
		for _, otherUserID := range userList {
			if otherUserID == userId {
				amountRespective = append(amountRespective, math.SmallestNonzeroFloat64)
				continue
			}
			amountRespective = append(amountRespective, amountPaid/float64(numberOfUsers))
		}
		return adjustExpenseRespective(userWhoPaid, userList, userId, expenseName, amountRespective)
	case "EXACT":
		var sumExpense float64
		for _, expense := range optionalIntSlice[0] {
			sumExpense += expense
		}
		if sumExpense < amountPaid {
			return errors.New(ErrSumSharesNotValid)
		}
		var amountRespective []float64
		for ind, otherUserID := range userList {
			if otherUserID == userId {
				amountRespective = append(amountRespective, math.SmallestNonzeroFloat64)
				continue
			}
			amountRespective = append(amountRespective, optionalIntSlice[0][ind])
		}
		return adjustExpenseRespective(userWhoPaid, userList, userId, expenseName, amountRespective)

	case "PERCENT":
		var sumPercentage float64
		for _, percentage := range optionalIntSlice[0] {
			sumPercentage += percentage
		}
		if int(sumPercentage) < 100 {
			return errors.New(ErrPercentageNotValid)
		}
		var amountRespective []float64
		for ind, otherUserID := range userList {
			if otherUserID == userId {
				amountRespective = append(amountRespective, math.SmallestNonzeroFloat64)
				continue
			}
			amountRespective = append(amountRespective, float64(optionalIntSlice[0][ind]*amountPaid)/float64(100))
		}
		return adjustExpenseRespective(userWhoPaid, userList, userId, expenseName, amountRespective)

	case "SHARE":
		var sumShare float64
		for _, share := range optionalIntSlice[0] {
			sumShare += share
		}
		eachShare := amountPaid / sumShare
		var amountRespective []float64
		for ind, otherUserID := range userList {
			if otherUserID == userId {
				amountRespective = append(amountRespective, math.SmallestNonzeroFloat64)
				continue
			}
			amountRespective = append(amountRespective, optionalIntSlice[0][ind]*eachShare)
		}
		return adjustExpenseRespective(userWhoPaid, userList, userId, expenseName, amountRespective)
	}
	return nil
}

func ShowBalance(userId string) error {
	log.Printf("userID %s balances:", userId)
	var currentUser *User
	if val, ok := usersMapping[userId]; ok {
		currentUser = val
	} else {
		return errors.New("map empty")
	}
	paymentList := currentUser.userPaymentList
	simplifiedPaymentList := currentUser.simplifiedUserPaymentList

	count := 0
	for user, amount := range paymentList {
		// u4 -> [(u1, 230), (u2, 240), (u3, 240)]
		// u1 -> [(u2, 250), (u3, 250)]
		// u2 -> [(u1, 250), (u4, 240)]
		if amount <= 0 {
			count++
			if count == len(paymentList) {
				log.Println("No Balance")
				break
			}
			continue
		}
		roundedAmount := fmt.Sprintf("%.2f", amount)
		log.Printf("user %s owes user %s : amount %s", user, userId, roundedAmount)
	}
	for user, amount := range paymentList {
		// u4 -> [(u1, 230), (u2, 240), (u3, 240)]
		// u1 -> [(u2, 250), (u3, 250)]
		// u2 -> [(u1, 250), (u4, 240)]
		if amount < 0 {
			continue
		}
		simplifiedPaymentList[user] = amount
	}
	//for user, amount := range simplifiedPaymentList {
	//	if amount < 0 {
	//		continue
	//	}
	//	roundedAmount := fmt.Sprintf("%.2f", amount)
	//	log.Printf("user %s owes user %s : amount %s and", userId, user, roundedAmount)
	//}
	return nil
}

func ShowBalanceAll() {
	for userId := range usersMapping {
		ShowBalance(userId)
	}
}

func ShowExpensesByName(expenseName string) error {
	if expenseData, ok := expenseNameMapping[expenseName]; ok {
		fmt.Printf("[%s -> ", expenseName)
		var maxUser string
		maxAmount := 0.0
		for user, amount := range expenseData {
			if amount > maxAmount {
				maxAmount = amount
				maxUser = user
			}
		}

		var users []string
		for user := range expenseData {
			users = append(users, user)
		}
		sort.Slice(users, func(i, j int) bool {
			return users[i] == maxUser
		})

		for _, user := range users {
			amount := expenseData[user]
			if user == maxUser {
				fmt.Printf("user %s paid amount %.0f, ", user, amount)
			} else {
				fmt.Printf("user %s owes amount %.0f ", user, amount)
			}
		}
		fmt.Println("]")
	} else {
		fmt.Println("Expense name not found.")
		return errors.New(ErrExpenseNameNotFound)
	}
	return nil
}

func ShowPassBookUserID(userId string) error {
	for _, expenseName := range allExpenseList {
		if expenseData, ok := expenseNameMapping[expenseName]; ok {
			var maxUser string
			maxAmount := 0.0
			for user, amount := range expenseData {
				if amount > maxAmount {
					maxAmount = amount
					maxUser = user
				}
			}

			var users []string
			for user := range expenseData {
				users = append(users, user)
			}
			sort.Slice(users, func(i, j int) bool {
				return users[i] == maxUser
			})

			var userExpenseExist bool
			for _, user := range users {
				if user == userId {
					userExpenseExist = true
					break
				}
			}
			if userExpenseExist == false {
				continue
			}
			fmt.Printf("[%s -> ", expenseName)
			for _, user := range users {
				amount := expenseData[user]
				if user == maxUser {
					fmt.Printf("user %s paid amount %.0f, ", user, amount)
				} else {
					fmt.Printf("user %s owes amount %.0f ", user, amount)
				}
			}
			fmt.Println("]")
		} else {
			fmt.Println("Expense name not found.")
			return errors.New(ErrExpenseNameNotFound)
		}
	}
	return nil
}
