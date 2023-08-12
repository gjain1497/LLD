package main

type State interface { //if there is any other type(struct)
	//which implements these functions(i.e have structure
	//like (func (s *idleState)ClickOnInsertCoinButton() error))
	//then we can say that idleState struct implements State
	ClickOnInsertCoinButton() error
	ClickOnStartProductSelectionButton() error
	InsertCoin(coin Coin) error
	ChooseProduct(codeNumber int) error
	GetChange(returnChangeMoney int) (int, error)
	DispenseProduct(codeNumber int) (Item, error)
	RefundFullMoney() ([]Coin, error)
}
