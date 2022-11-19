package abstract_factory

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() string
}

type Shirt struct {
	Logo string
	size int
}

func (s *Shirt) setLogo(logo string){
	s.Logo = logo
}

func (s *Shirt) getLogo() string {
	return s.Logo
}

func (s *Shirt) setSize(size int) {
	s.Size = size
}

func (s *Shirt) getSize() int {
	return s.Size
}