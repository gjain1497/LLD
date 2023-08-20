package main

import (
	"fmt"
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

	err1 := sp.AddExpense("hotel", 1000, "u1", 4, []string{"u1", "u2", "u3", "u4"}, "EQUAL")
	if err1 != nil {
		log.Println(err1)
	}
	//u2 owes u1 owes 250 for hotel
	err2 := sp.AddExpense("breakfast", 1250, "u1", 2, []string{"u2", "u3"}, "EXACT", []float64{370, 880})
	if err2 != nil {
		log.Println(err2)
	}
	err3 := sp.AddExpense("flight", 1200, "u4", 4, []string{"u1", "u2", "u3", "u4"}, "PERCENT", []float64{40, 20, 20, 20})
	if err3 != nil {
		log.Println(err3)
	}
	err4 := sp.AddExpense("Dinner", 1200, "u4", 4, []string{"u1", "u2", "u3", "u4"}, "SHARE", []float64{2, 1, 1, 1})
	if err4 != nil {
		log.Println(err4)
	}
	//sp.ShowBalance("u1") //for specific user id
	//sp.ShowBalance("u2") //for specific user id
	//sp.ShowBalance("u3") //for specific user id
	//sp.ShowBalance("u4") //for specific user id

	fmt.Println()
	fmt.Println()
	fmt.Println("Showing balance for all:")
	sp.ShowBalanceAll() //for all
	fmt.Println()
	fmt.Println()
	fmt.Println("Showing expense by name:")
	sp.ShowExpensesByName("hotel")
	sp.ShowExpensesByName("breakfast")
	sp.ShowExpensesByName("flight")
	sp.ShowExpensesByName("Dinner")
	fmt.Println()
	fmt.Println()

	fmt.Println("Passbook for user u4:")
	sp.ShowExpensesByUserID("u4")

	//sp.ShowExpensesByName("flight")
}

//
//(Kulche -> [u1->[(u2,20),(u3,10),(u4,80)]]), (Breakfast -> [u1->[(u2,20),(u3,10),(u4,90)]])

// [kulche -> (u1, 1000), (u2, 250), (u3, 250), (u4, 250)]

//getExpense("kulcha")

//u1 -> u2 (200) u2-> u3 (400) u3->u4 (300)
// u1->u3 (200), u1->u2(200)
// u1->u4 (200)
