package main

import (
	sp "github.com/gjain1497/LLD/splitwise/splitwise"
	"log"
)

func main() {

	//u1 := sp.NewUser("u1")
	//u2 := sp.NewUser("u1")
	//u3 := sp.NewUser("u1")
	//u4 := sp.NewUser("u1")

	sp.AddUser("u1", "Girish Jain", 7888873022)
	sp.AddUser("u2", "Nipun Jain", 7888873022)
	sp.AddUser("u3", "Gagandeep Ahuja", 7888873022)
	sp.AddUser("u4", "Madhav Mehta ", 7888873022)

	err1 := sp.AddExpense(1000, "u1", 4, []string{"u1", "u2", "u3", "u4"}, "EQUAL")
	if err1 != nil {
		log.Println(err1)
	}
	err2 := sp.AddExpense(1250, "u1", 2, []string{"u2", "u3"}, "EXACT", []float64{370, 870})
	if err2 != nil {
		log.Println(err2)
	}
	err3 := sp.AddExpense(1200, "u4", 4, []string{"u1", "u2", "u3", "u4"}, "PERCENT", []float64{40, 20, 20, 20})
	if err3 != nil {
		log.Println(err3)
	}
	//sp.ShowBalance("u1") //for specific user id
	//sp.ShowBalance("u2") //for specific user id
	//sp.ShowBalance("u3") //for specific user id
	//sp.ShowBalance("u4") //for specific user id

	sp.ShowBalanceAll() //for all
}
