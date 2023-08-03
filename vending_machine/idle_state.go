package main

import (
	"errors"
	"fmt"
	"log"
)

type IdleState struct{}

func NewIdleState() *IdleState {
	fmt.Println("Currenlty machine is in IdleState")
	return &IdleState{}
}

func (s IdleState) ClickOnInsertCoinButton(machine *VendingMachine) error {
	machine.SetVendingMachineState(NewHasMoneyState())
	log.Println("Changing state from idle to next state")
	return nil
}
func (s IdleState) ClickOnStartProductSelectionButton(machine *VendingMachine) error {
	return errors.New("error")
}
func (s IdleState) InsertCoin(machine *VendingMachine, coin Coin) error {
	return errors.New("error")
}
func (s IdleState) ChooseProduct(machine *VendingMachine, codeNumber int) error {
	return errors.New("error")
}
func (s IdleState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("error")
}
func (s IdleState) DispenseProduct(machine *VendingMachine, codeNumber int) (Item, error) {
	return Item{}, errors.New("error")
}
func (s IdleState) RefundFullMoney(machine *VendingMachine) ([]Coin, error) {
	return nil, errors.New("error")
}
