package impl

import (
	"errors"
	"fmt"
)

type HasMoneyState struct{}

func (hms HasMoneyState) ClickOnInsertCoinButton(machine VendingMachine) error {
	return errors.New("you cannot click on the insert button in hasMoney state")
}

func (hms HasMoneyState) GetChange(returnChangeMoney int) (int, error) {
	return 0, errors.New("you cannot get change in hasMoney state")
}

func (hms HasMoneyState) UpdateInventory(machine VendingMachine, item Item, codeNumber int) error {
	return errors.New("you cannot update inventory in hasMoney state")
}

func (hms HasMoneyState) RefundFullMoney(machine VendingMachine) ([]Coin, error) {
	return nil, errors.New("you cannot refund money in hasMoney state")
}

func (hms HasMoneyState) DispenseProduct(machine VendingMachine, codeNumber int) (Item, error) {
	return nil, errors.New("product cannot be dispensed in hasMoney state")
}

func (hms HasMoneyState) ChooseProduct(machine VendingMachine, codeNumber int) error {
	return errors.New("you need to click on the start product selection button first")
}

func (hms HasMoneyState) InsertCoin(machine VendingMachine, coin Coin) error {
	fmt.Println("Accepted the coin")
	machine.GetCoinList().Add(coin)
	return nil
}

func (hms HasMoneyState) ClickOnStartProductSelectionButton(machine VendingMachine) error {
	machine.SetVendingMachineState(vendingmachine.NewSelectionState())
	return nil
}
