package abstract_factory

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type Concrete struct {
	shoe Shoe
}

func NewShoe() IProduct {

	return Concrete{
		shoe: Shoe{"ciquan", 0},
	}
}