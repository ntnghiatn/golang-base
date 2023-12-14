package abstract_factory

type ConcreteWinterCombo struct {
	//product map[string]interface{}
	shoe  Shoe
	short Short
}

func (c ConcreteWinterCombo) GetShoe() Shoe {
	//TODO implement me
	panic("implement me")
	return c.shoe
}

func (c ConcreteWinterCombo) GetShort() Short {
	//TODO implement me
	return c.short
}
