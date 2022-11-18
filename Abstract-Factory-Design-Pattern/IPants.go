package abstract_factory

type IPants interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() string
}

type Pants struct {
	Logo string
	size int
}

func (s *Pants) setLogo(logo string){
	s.Logo = logo
}

func (s *Pants) getLogo() string {
	return s.Logo
}

func (s *Pants) setSize(size int) {
	s.Size = size
}

func (s *Pants) getSize() int {
	return s.Size
}