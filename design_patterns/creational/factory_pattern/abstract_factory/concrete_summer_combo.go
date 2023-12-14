package abstract_factory

import "fmt"

type ConcreteSummerCombo struct{}

func (c ConcreteSummerCombo) GetShoe() shoe {
	//TODO implement me
	fmt.Println("Dep Lao mua he")
	return shoe{"SSSS", 1}
}

func (c ConcreteSummerCombo) GetShort() Short {
	//TODO implement me
	fmt.Println("Tat luoi cho mat")
	return Short{}
}
