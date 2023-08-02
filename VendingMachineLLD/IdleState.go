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

func (s *IdleState) ClickOnInsertCoinButton(machine VendingMachine) {
	log.Println("Changing state from idle to next state")
}
func (s *IdleState) ClickOnStartProductSelectionButton() error {
	return errors.New("error")
}
func (s *IdleState) InsertCoin(coin Coin) error {
	return errors.New("error")
}
func (s *IdleState) ChooseProduct(codeNumber int) error {
	return errors.New("error")
}
func (s *IdleState) GetChange(returnChangeMoney int) error {
	return errors.New("error")
}
func (s *IdleState) DispenseProduct(codeNumber int) error {
	return errors.New("error")
}
func (s *IdleState) RefundFullMoney() ([]Coin, error) error{
	return errors.New("error")
}
func (s *IdleState) UpdateInventory(item Item, codeNumber int) error {
	return errors.New("error")
}
