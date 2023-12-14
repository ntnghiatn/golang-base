package abstract_factory

type IShort interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type ConcreteShort struct {
	short Short
}

// NewConcreteShort - Factory method
func NewConcreteShort() IProduct {
	return &ConcreteShort{
		short: Short{"nameShort", 0},
	}
}
func (c *ConcreteShort) setLogo(logo string) {
	//TODO implement me
	panic("implement me")
}

func (c *ConcreteShort) setSize(size int) {
	//TODO implement me
	panic("implement me")
}

func (c *ConcreteShort) GetLogo() string {
	//TODO implement me
	panic("implement me")
}

func (c *ConcreteShort) GetSize() int {
	//TODO implement me
	panic("implement me")
}
