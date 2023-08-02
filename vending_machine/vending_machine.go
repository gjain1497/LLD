package main

type VendingMachine struct {
	vendingMachineState State
	inventory           Inventory
	coinList            []Coin
}

func NewVendingMachine() VendingMachine {
	return &VendingMachine{
		vendingMachineState: vendingstates.NewIdleState(),
		inventory:           NewInventory(10),
		coinList:            make([]Coin, 0),
	}
}

// Getter methods
func (vm VendingMachine) GetVendingMachineState() State {
	return vm.vendingMachineState
}

func (vm VendingMachine) GetInventory() Inventory {
	return vm.inventory
}

func (vm VendingMachine) GetCoinList() []Coin {
	return vm.coinList
}

// Setter methods

func (vm VendingMachine) SetVendingMachineState(state State) {
	vm.vendingMachineState = state
}

func (vm VendingMachine) SetInventory(inventory Inventory) {
	vm.inventory = inventory
}

func (vm VendingMachine) SetCoinList(coinList []Coin) {
	vm.coinList = coinList
}
