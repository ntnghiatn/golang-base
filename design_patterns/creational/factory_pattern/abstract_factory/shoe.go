package abstract_factory

type shoe struct {
	logo string
	size int
}

func (s shoe) GetLogo() string {
	return s.logo
}

func (s shoe) GetSize() int {
	return s.size
}
