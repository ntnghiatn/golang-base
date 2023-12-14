package abstract_factory

type IProduct interface {
	IShort
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}
