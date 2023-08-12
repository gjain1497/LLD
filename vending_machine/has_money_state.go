package main

import (
	"errors"
	"fmt"
	"log"
)

type HasMoneyState struct {
	machine *VendingMachine
}

func NewHasMoneyState() *HasMoneyState {
	fmt.Println("Currenlty machine is in HasMoneyState")
	return &HasMoneyState{}
}

func (hms HasMoneyState) ClickOnInsertCoinButton() error {
	return errors.New("you cannot click on the insert button in hasMoney state")
}

func (hms HasMoneyState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("you cannot get change in hasMoney state")
}

func (hms HasMoneyState) RefundFullMoney() ([]Coin, error) {
	return nil, errors.New("you cannot refund money in hasMoney state")
}

func (hms HasMoneyState) DispenseProduct(codeNumber int) (Item, error) {
	return Item{}, errors.New("product cannot be dispensed in hasMoney state")
}

func (hms HasMoneyState) ChooseProduct(codeNumber int) error {
	return errors.New("you need to click on the start product selection button first")
}

func (hms HasMoneyState) InsertCoin(coin Coin) error {
	fmt.Println("Accepted the coin")
	coinsArray := hms.machine.GetCoinList()
	coinsArray = append(coinsArray, coin)
	for _, val := range coinsArray {
		log.Println("coin[i] ", val)
	}
	return nil
}

func (hms HasMoneyState) ClickOnStartProductSelectionButton() error {
	hms.machine.currentState = NewSelectionState()
	return nil
}
