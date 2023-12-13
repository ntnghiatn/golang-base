package product

type IShort interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type ConcreteShort struct {
	short Short
}

func NewShort() IProduct {
	return &ConcreteShort{
		short: Short{},
	}
}
