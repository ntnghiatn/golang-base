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

// NewShort - Factory method
func NewShort() IProduct {
	return &ConcreteShort{
		short: Short{},
	}
}
