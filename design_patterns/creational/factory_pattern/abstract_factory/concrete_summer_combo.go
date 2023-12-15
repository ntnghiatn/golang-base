package abstract_factory

type ConcreteSummerCombo struct {
	shoe  shoe
	short Short
}

func NewSummerCombo() ICombo {
	return ConcreteSummerCombo{shoe: shoe{
		logo: "DEP LAO",
		size: 22,
	}, short: Short{name: "Vo Luoi", size: 222}}
}
func (c ConcreteSummerCombo) GetShoe() shoe {
	return c.shoe
}

func (c ConcreteSummerCombo) GetShort() Short {
	return c.short
}
