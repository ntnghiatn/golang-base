package abstract_factory

type ConcreteShoe struct {
	product shoe
}

func NewConcreteShoe() IProduct {
	return &ConcreteShoe{
		product: shoe{"GiayRach", 0},
	}
}

func NewShoeForWinter() IProduct {
	return &ConcreteShoe{
		product: shoe{"Dong", 1},
	}
}

func (c *ConcreteShoe) setLogo(logo string) {
	c.product.logo = logo
}

func (c *ConcreteShoe) setSize(size int) {
	c.product.size = size
}

func (c *ConcreteShoe) GetLogo() string {
	return c.product.logo
}

func (c *ConcreteShoe) GetSize() int {
	return c.product.size
}
