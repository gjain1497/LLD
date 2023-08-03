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

	vendingMachineState := vendingMachine.GetVendingMachineState()
	err := vendingMachineState.ClickOnInsertCoinButton(vendingMachine)
	if err != nil {
		fmt.Println(err)
	}

	vendingMachineState = vendingMachine.GetVendingMachineState()
	err = vendingMachineState.InsertCoin(vendingMachine, NICKEL)
	if err != nil {
		fmt.Println(err)
	}
	err = vendingMachineState.InsertCoin(vendingMachine, QUARTER)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("|")
	fmt.Println("Clicking on the ProductSelectionButton")
	fmt.Println("|")
	err = vendingMachineState.ClickOnStartProductSelectionButton(vendingMachine)
	if err != nil {
		fmt.Println(err)
	}

	vendingMachineState = vendingMachine.GetVendingMachineState()
	err = vendingMachineState.ChooseProduct(vendingMachine, 102)
	if err != nil {
		fmt.Println(err)
	}

	displayInventory(vendingMachine)
}

func displayInventory(vendingMachine *VendingMachine) {
	slots := vendingMachine.GetInventory().GetItemShelf()
	for i := 0; i < len(slots); i++ {
		fmt.Printf("CodeNumber: %d Item: %s Price: %d isAvailable: %t\n",
			slots[i].GetCode(),
			slots[i].GetItem().GetType(),
			slots[i].GetItem().GetPrice(),
			!slots[i].IsSoldOut())
	}
}

func fillUpTheInventory(vendingMachine *VendingMachine) {
	slots := vendingMachine.GetInventory().GetItemShelf()
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
