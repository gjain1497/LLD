package main

import (
	"errors"
	"fmt"
)

type DispenseState struct{}

func NewDispenseState(machine *VendingMachine, codeNumber int) *DispenseState {
	fmt.Println("Currently Vending Machine is in DispenseState")
	machine.GetVendingMachineState().DispenseProduct(machine, codeNumber)
	return &DispenseState{}
}

func (ds DispenseState) ClickOnInsertCoinButton(machine *VendingMachine) error {
	return errors.New("insert coin button cannot be clicked on Dispense state")
}

func (ds DispenseState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("change cannot be returned in Dispense state")
}

func (ds DispenseState) RefundFullMoney(machine *VendingMachine) ([]Coin, error) {
	return nil, errors.New("refund cannot happen in Dispense state")
}

func (ds DispenseState) DispenseProduct(machine *VendingMachine, codeNumber int) (Item, error) {
	fmt.Println("Product has been dispensed")
	inv := machine.GetInventory()
	item, err := inv.GetItem(codeNumber)
	if err != nil {
		return Item{}, err
	}
	inv.UpdateSoldOutItem(codeNumber)
	machine.SetVendingMachineState(NewIdleState())
	return item, nil
}

func (ds DispenseState) ChooseProduct(machine *VendingMachine, codeNumber int) error {
	return errors.New("product cannot be chosen in Dispense state")
}

func (ds DispenseState) InsertCoin(machine *VendingMachine, coin Coin) error {
	return errors.New("coin cannot be inserted in Dispense state")
}

func (ds DispenseState) ClickOnStartProductSelectionButton(machine *VendingMachine) error {
	return errors.New("product selection button cannot be clicked in Dispense state")
}
