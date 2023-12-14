package abstract_factory

type Shoe struct {
	logo string
	size int
}

func (c Concrete) setLogo(logo string) {
	//TODO implement me
	panic("implement set logo")
}

func (c Concrete) setSize(size int) {
	//TODO implement me
	panic("implement set size")
}

func (c Concrete) GetLogo() string {
	//TODO implement me
	return c.shoe.logo
}

func (c Concrete) GetSize() int {
	//TODO implement me
	return c.shoe.size
}
