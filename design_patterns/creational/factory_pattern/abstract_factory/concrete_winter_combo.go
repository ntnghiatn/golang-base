package abstract_factory

type ConcreteWinterCombo struct {
	//product map[string]interface{}
	shoe  shoe
	short Short
}

func NewWinterCombo() ICombo {
	return &ConcreteWinterCombo{short: Short{name: "DDD", size: 1}, shoe: shoe{logo: "Win Nike", size: 0}}
}

func (c ConcreteWinterCombo) GetShoe() shoe {
	return c.shoe
}

func (c ConcreteWinterCombo) GetShort() Short {
	return c.short
}
