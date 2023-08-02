package impl

import (
	"errors"
	"fmt"
)

type DispenseState struct {
}

func NewDispenseState(machine VendingMachine, codeNumber int) (DispenseState, error) {
	fmt.Println("Currently Vending Machine is in DispenseState")
	item, err := machine.GetInventory().DispenseProduct(machine, codeNumber)
	if err != nil {
		return nil, err
	}
	machine.GetInventory().UpdateSoldOutItem(codeNumber)
	machine.SetVendingMachineState(vendingmachine.NewIdleState(machine))
	return &DispenseState{}, nil
}

func (ds DispenseState) ClickOnInsertCoinButton(machine *vendingmachine.VendingMachine) error {
	return errors.New("insert coin button cannot be clicked on Dispense state")
}

func (ds DispenseState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("change cannot be returned in Dispense state")
}

func (ds DispenseState) RefundFullMoney(machine VendingMachine) ([]Coin, error) {
	return nil, errors.New("refund cannot happen in Dispense state")
}

func (ds DispenseState) UpdateInventory(machine VendingMachine, item Item, codeNumber int) error {
	return errors.New("inventory cannot be updated in Dispense state")
}

func (ds DispenseState) DispenseProduct(machine VendingMachine, codeNumber int) (Item, error) {
	fmt.Println("Product has been dispensed")
	item, err := machine.GetInventory().GetItem(codeNumber)
	if err != nil {
		return nil, err
	}
	machine.GetInventory().UpdateSoldOutItem(codeNumber)
	machine.SetVendingMachineState(vendingmachine.NewIdleState(machine))
	return item, nil
}

func (ds DispenseState) ChooseProduct(machine VendingMachine, codeNumber int) error {
	return errors.New("product cannot be chosen in Dispense state")
}

func (ds DispenseState) InsertCoin(machine VendingMachine, coin Coin) error {
	return errors.New("coin cannot be inserted in Dispense state")
}

func (ds DispenseState) ClickOnStartProductSelectionButton(machine VendingMachine) error {
	return errors.New("product selection button cannot be clicked in Dispense state")
}
