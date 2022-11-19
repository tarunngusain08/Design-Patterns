package abstract_factory

type PeterEngland struct {
}

type PeterEnglandShirt struct {
	Shirt
}

type PeterEnglandPants struct {
	Pants
}

type PeterEnglandBlazer struct {
	Blazer
}

func (p *PeterEngland) makeShirt() IShirt{
	return &PeterEnglandShirt{
		Shirt : Shirt{
			logo: "PeterEngland",
			size: 42
		}
	}
}

func (p	*PeterEngland) makePants() IPants{
	return &PeterEnglandPants{
		Pants : Pants{
			logo: "PeterEngland",
			size: 30
		}
	}
}

func (p *PeterEngland) makeBlazer() IBlazer{
	return &PeterEnglandBlazer{
		Blazer : Blazer{
			logo: "PeterEngland",
			size: 44
		}
	}
}