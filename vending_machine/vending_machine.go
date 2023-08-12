package main

type VendingMachine struct {
	currentState State
	inventory    Inventory
	coinList     []Coin
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		currentState: NewIdleState(),
		inventory:    NewInventory(10),
		coinList:     make([]Coin, 0),
	}
}

func (v *VendingMachine) GetVendingMachineState() State {
	return v.currentState
}

func (v *VendingMachine) GetInventory() Inventory {
	return v.inventory
}

func (v *VendingMachine) GetCoinList() []Coin {
	return v.coinList
}

func (v *VendingMachine) ClickOnInsertCoinButton() error {
	return v.currentState.ClickOnInsertCoinButton()
}
func (v *VendingMachine) ClickOnStartProductSelectionButton() error {
	return v.currentState.ClickOnStartProductSelectionButton()
}
func (v *VendingMachine) InsertCoin(coin Coin) error {
	return v.currentState.InsertCoin(coin)
}
func (v *VendingMachine) ChooseProduct(codeNumber int) error {
	return v.currentState.ChooseProduct(codeNumber)
}

//func (v *VendingMachine) DispenseProduct(machine *VendingMachine, codeNumber int) (Item, error) {
//
//}
//func (v *VendingMachine) RefundFullMoney(machine *VendingMachine) ([]Coin, error) {
//
//}
