package abstract_factory

type Levis struct {
}

type LevisShirt struct {
	Shirt
}

type LevisPants struct {
	Pants
}

type LevisBlazer struct {
	Blazer
}

func (p *Levis) makeShirt() IShirt{
	return &LevisShirt{
		Shirt: Shirt{
			logo: "Levis",
			size: 44,
		}
	}
}

func (p	*Levis) makePants() IPants{
	return &LevisPants{
		Pants : Pants{
			logo: "Levis",
			size: 30
		}
	}
}
func (p *Levis) makeBlazer() IBlazer{
	return &levisBlazer{
		Blazer : Blazer{
			logo: "Levis",
			size: 44
		}
	}
}