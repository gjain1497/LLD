package main

import (
	"errors"
	"fmt"
	"log"
)

type SelectionState struct {
	machine *VendingMachine
}

func NewSelectionState() *SelectionState {
	fmt.Println("Currenlty machine is in SelectionState")
	return &SelectionState{}
}

func (ss SelectionState) ClickOnInsertCoinButton() error {
	return errors.New("you cannot click on the insert coin button in Selection state")
}

func (ss SelectionState) GetChange(returnChangeMoney int) (int, error) {
	// Actual logic should be to return COINS in the dispense tray,
	// but for simplicity, I am just returning the amount to be refunded
	fmt.Println("Returned the changes in the coin Dispense Tray:", returnChangeMoney)
	return returnChangeMoney, nil
}

func (ss SelectionState) RefundFullMoney() ([]Coin, error) {
	fmt.Println("Returned the change in Coin Dispense Tray")
	ss.machine.currentState = NewIdleState()
	return ss.machine.GetCoinList(), nil
}

func (ss SelectionState) DispenseProduct(codeNumber int) (Item, error) {
	return Item{}, errors.New("product cannot be dispensed in Selection state")
}

func (ss SelectionState) ChooseProduct(codeNumber int) error {
	// 1. Get the item of this code number
	log.Println("ChooseProduct func")
	inv := ss.machine.GetInventory()
	item, err := inv.GetItem(codeNumber)
	if err != nil {
		return err
	}

	// 2. Total amount paid by User
	paidByUser := 0
	log.Println("paidByUser ", paidByUser)

	for _, coin := range ss.machine.GetCoinList() {
		log.Println("coin Value ", coin.Value)
		paidByUser += coin.Value
	}
	log.Println("paidByUser after", paidByUser)

	// 3. Compare product price and amount paid by the user
	if paidByUser < item.Price {
		fmt.Printf("Insufficient amount, the product you selected is for the price: %d and you paid: %d\n", item.Price, paidByUser)
		_, err := ss.RefundFullMoney()
		if err != nil {
			return err
		}
		return errors.New("insufficient amount")
	} else if paidByUser >= item.Price {
		if paidByUser > item.Price {
			ss.GetChange(paidByUser - item.Price)
		}
		ss.machine.currentState = NewDispenseState(ss.machine, codeNumber)
	}
	return nil
}

func (ss SelectionState) InsertCoin(coin Coin) error {
	return errors.New("you cannot insert a coin in the selection state")
}

func (ss SelectionState) ClickOnStartProductSelectionButton() error {
	return errors.New("you cannot insert a coin in the selection state")
}
