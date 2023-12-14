package abstract_factory

import "fmt"

type ISummerCombo interface {
	GetShoe()
	GetShort()
}

type ConcreteSummerCombo struct{}

func (c ConcreteSummerCombo) GetShoe() {
	//TODO implement me
	fmt.Println("Dep Lao mua he")
}

func (c ConcreteSummerCombo) GetShort() {
	//TODO implement me
	fmt.Println("Tat luoi cho mat")
}
