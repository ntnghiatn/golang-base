package abstract_factory

import "fmt"

type IwinterCombo interface {
	GetShoe()
	GetShort()
}

type ConcreteWinterCombo struct{}

func (c ConcreteWinterCombo) GetShoe() {
	//TODO implement me
	fmt.Println("Giày cao cổ for mùa pđông ")
}

func (c ConcreteWinterCombo) GetShort() {
	//TODO implement me
	fmt.Println("tat cao tren dau gau")
}
