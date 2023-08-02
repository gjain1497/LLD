package main

type VendingMachine struct {
	vendingMachineState State
	inventory           Inventory
	coinList            []Coin
}

// Getter methods

func (vm *VendingMachine) GetVendingMachineState() State {
	return vm.vendingMachineState
}

func (vm *VendingMachine) GetInventory() Inventory {
	return vm.inventory
}

func (vm *VendingMachine) GetCoinList() []Coin {
	return vm.coinList
}

// Setter methods

func (vm *VendingMachine) SetVendingMachineState(state State) {
	vm.vendingMachineState = state
}

func (vm *VendingMachine) SetInventory(inventory Inventory) {
	vm.inventory = inventory
}

func (vm *VendingMachine) SetCoinList(coinList []Coin) {
	vm.coinList = coinList
}
