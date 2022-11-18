package abstract_factory

type IBlazer interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() string
}

type Blazer struct {
	Logo string
	size int
}

func (s *Blazer) setLogo(logo string){
	s.Logo = logo
}

func (s *Blazer) getLogo() string {
	return s.Logo
}

func (s *Blazer) setSize(size int) {
	s.Size = size
}

func (s *Blazer) getSize() int {
	return s.Size
}