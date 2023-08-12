package main

import (
	"errors"
	"fmt"
)

type DispenseState struct {
	machine *VendingMachine
}

func NewDispenseState(machine *VendingMachine, codeNumber int) *DispenseState {
	fmt.Println("Currently Vending Machine is in DispenseState")
	machine.GetVendingMachineState().DispenseProduct(codeNumber)
	return &DispenseState{}
}

func (ds DispenseState) ClickOnInsertCoinButton() error {
	return errors.New("insert coin button cannot be clicked on Dispense state")
}

func (ds DispenseState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("change cannot be returned in Dispense state")
}

func (ds DispenseState) RefundFullMoney() ([]Coin, error) {
	return nil, errors.New("refund cannot happen in Dispense state")
}

func (ds DispenseState) DispenseProduct(codeNumber int) (Item, error) {
	fmt.Println("Product has been dispensed")
	inv := ds.machine.GetInventory()
	item, err := inv.GetItem(codeNumber)
	if err != nil {
		return Item{}, err
	}
	inv.UpdateSoldOutItem(codeNumber)
	ds.machine.currentState = NewIdleState()
	return item, nil
}

func (ds DispenseState) ChooseProduct(codeNumber int) error {
	return errors.New("product cannot be chosen in Dispense state")
}

func (ds DispenseState) InsertCoin(coin Coin) error {
	return errors.New("coin cannot be inserted in Dispense state")
}

func (ds DispenseState) ClickOnStartProductSelectionButton() error {
	return errors.New("product selection button cannot be clicked in Dispense state")
}
