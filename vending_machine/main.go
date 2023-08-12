package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := NewVendingMachine()

	fmt.Println("|")
	fmt.Println("Filling up the inventory")
	fmt.Println("|")

	fillUpTheInventory(vendingMachine)
	displayInventory(vendingMachine)

	fmt.Println("|")
	fmt.Println("Clicking on the insert coin button")
	fmt.Println("|")

	//Here I am first getting state from vending machine state and then calling ClickOnInsertCoinButton
	//but rather I should directly call vendingMachine.ClickOnInsertCoinButton
	//So Gagan code is providing me that option, in which we are putting the below
	//code functions inside vending machine file

	//vendingMachineState := vendingMachine.GetVendingMachineState()
	//err := vendingMachineState.ClickOnInsertCoinButton(vendingMachine)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//so below is a better practice
	//and if we take vendingMachine as param
	if err := vendingMachine.ClickOnInsertCoinButton(); err != nil {
		fmt.Println("Err inserting first time ")
	}

	//vendingMachineState = vendingMachine.GetVendingMachineState() ---> on same grounds commented this line as well
	if err := vendingMachine.InsertCoin(NICKEL); err != nil {
		fmt.Println("Err inserting first time")
	}
	if err := vendingMachine.InsertCoin(QUARTER); err != nil {
		fmt.Println("Err inserting second coin ")
	}

	fmt.Println("|")
	fmt.Println("Clicking on the ProductSelectionButton")
	fmt.Println("|")
	//err = vendingMachineState.ClickOnStartProductSelectionButton(vendingMachine) ---> on same grounds commented this line as well
	if err := vendingMachine.ClickOnStartProductSelectionButton(); err != nil {
		fmt.Println("Err clicking on product selection button")
	}

	//vendingMachineState = vendingMachine.GetVendingMachineState()
	//err = vendingMachineState.ChooseProduct(vendingMachine, 102)
	//if err != nil {
	//	fmt.Println(err)
	//}

	if err := vendingMachine.ChooseProduct(102); err != nil {
		fmt.Println("Err clicking on product selection button")
	}

	displayInventory(vendingMachine)
}

func displayInventory(vendingMachine *VendingMachine) {
	inv := vendingMachine.GetInventory()
	slots := inv.GetItemShelf()
	for i := 0; i < len(slots); i++ {
		fmt.Printf("CodeNumber: %d Item: %s Price: %d isAvailable: %t\n",
			slots[i].GetCode(),
			slots[i].GetItem().GetType(),
			slots[i].GetItem().GetPrice(),
			!slots[i].IsSoldOut())
	}
}

func fillUpTheInventory(vendingMachine *VendingMachine) {
	inv := vendingMachine.GetInventory()
	slots := inv.GetItemShelf()
	for i := 0; i < len(slots); i++ {
		newItem := &Item{}
		if i >= 0 && i < 3 {
			newItem.SetType(COKE)
			newItem.SetPrice(12)
		} else if i >= 3 && i < 5 {
			newItem.SetType(PEPSI)
			newItem.SetPrice(9)
		} else if i >= 5 && i < 7 {
			newItem.SetType(JUICE)
			newItem.SetPrice(13)
		} else if i >= 7 && i < 10 {
			newItem.SetType(SODA)
			newItem.SetPrice(7)
		}
		log.Println("newItem: ", newItem)
		slots[i].SetItem(*newItem)
		slots[i].SetSoldOut(false)

		log.Println("slots[i] ", slots[i])
	}
}
