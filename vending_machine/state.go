package main

type State interface { //if there is any other type(struct)
	//which implements these functions(i.e have structure
	//like (func (s *idleState)ClickOnInsertCoinButton() error))
	//then we can say that idleState struct implements State
	ClickOnInsertCoinButton(machine *VendingMachine) error
	ClickOnStartProductSelectionButton(machine *VendingMachine) error
	InsertCoin(machine *VendingMachine, coin Coin) error
	ChooseProduct(machine *VendingMachine, codeNumber int) error
	GetChange(returnChangeMoney int) (int, error)
	DispenseProduct(machine *VendingMachine, codeNumber int) (Item, error)
	RefundFullMoney(machine *VendingMachine) ([]Coin, error)
}
