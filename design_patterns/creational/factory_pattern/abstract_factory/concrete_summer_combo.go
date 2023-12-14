package abstract_factory

import "fmt"

type ConcreteSummerCombo struct{}

func (c ConcreteSummerCombo) GetShoe() Shoe {
	//TODO implement me
	fmt.Println("Dep Lao mua he")
	return Shoe{"SSSS", 1}
}

func (c ConcreteSummerCombo) GetShort() Short {
	//TODO implement me
	fmt.Println("Tat luoi cho mat")
	return Short{}
}
