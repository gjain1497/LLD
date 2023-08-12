package main

import (
	"errors"
	"fmt"
	"log"
)

type IdleState struct {
	machine *VendingMachine
}

func NewIdleState() *IdleState {
	fmt.Println("Currenlty machine is in IdleState")
	return &IdleState{}
}

func (s IdleState) ClickOnInsertCoinButton() error {
	s.machine.currentState = NewHasMoneyState()
	log.Println("Changing state from idle to next state")
	return nil
}
func (s IdleState) ClickOnStartProductSelectionButton() error {
	return errors.New("error")
}
func (s IdleState) InsertCoin(coin Coin) error {
	return errors.New("error")
}
func (s IdleState) ChooseProduct(codeNumber int) error {
	return errors.New("error")
}
func (s IdleState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("error")
}
func (s IdleState) DispenseProduct(codeNumber int) (Item, error) {
	return Item{}, errors.New("error")
}
func (s IdleState) RefundFullMoney() ([]Coin, error) {
	return nil, errors.New("error")
}
