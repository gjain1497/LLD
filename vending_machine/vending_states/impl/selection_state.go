package impl

import (
	"errors"
	"fmt"
)

type SelectionState struct {
}

func (ss SelectionState) ClickOnInsertxCoinButton(machine VendingMachine) error {
	return errors.New("you cannot click on the insert coin button in Selection state")
}

func (ss SelectionState) GetChange(returnChangeMoney int) int {
	// Actual logic should be to return COINS in the dispense tray,
	// but for simplicity, I am just returning the amount to be refunded
	fmt.Println("Returned the changes in the coin Dispense Tray:", returnChangeMoney)
	return returnChangeMoney
}

func (ss SelectionState) UpdateInventory(machine VendingMachine, item Item, codeNumber int) error {
	return errors.New("Inventory cannot be updated in Selection state")
}

func (ss SelectionState) RefundFullMoney(machine VendingMachine) ([]Coin, error) {
	fmt.Println("Returned the change in Coin Dispense Tray")
	machine.SetVendingMachineState(vendingmachine.NewIdleState(machine))
	return machine.GetCoinList(), nil
}

func (ss SelectionState) DispenseProduct(machine VendingMachine, codeNumber int) (Item, error) {
	return vendingmachine.Item{}, errors.New("product cannot be dispensed in Selection state")
}

func (ss SelectionState) ChooseProduct(machine VendingMachine, codeNumber int) error {
	// 1. Get the item of this code number
	item, err := machine.GetInventory().GetItem(codeNumber)
	if err != nil {
		return err
	}

	// 2. Total amount paid by User
	paidByUser := 0
	for _, coin := range machine.GetCoinList() {
		paidByUser += coin.Value
	}

	// 3. Compare product price and amount paid by the user
	if paidByUser < item.Price {
		fmt.Printf("Insufficient amount, the product you selected is for the price: %d and you paid: %d\n", item.Price, paidByUser)
		_, err := ss.RefundFullMoney(machine)
		if err != nil {
			return err
		}
		return errors.New("insufficient amount")
	} else if paidByUser >= item.Price {
		if paidByUser > item.Price {
			ss.GetChange(paidByUser - item.Price)
		}
		machine.SetVendingMachineState(vendingmachine.NewDispenseState(machine, codeNumber))
	}
	return nil
}

func (ss SelectionState) InsertCoin(machine VendingMachine, coin Coin) error {
	return errors.New("you cannot insert a coin in the selection state")
}

func (ss SelectionState) ClickOnStartProductSelectionButton(machine VendingMachine) {
}
