package abstract_factory

type PeterEngland struct {
}

func (p *PeterEngland) makeShirt() IShirt{}
func (p	*PeterEngland) makePants() IPants{}
func (p *PeterEngland) makeBlazer() IBlazer{}