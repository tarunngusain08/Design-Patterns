package abstract_factory

type ISuitFactory interface {
	makeShirt() IShirt
	makePants() IPants
	makeBlazer() IBlazer
}

func GetSuitFactory(brand string) (ISuitFactory,error) {
	if brand == "PeterEngland"{
		return &PeterEngland{},nil
	}
	if brand == "Levis"{
		return &Levis{},nil
	}
	return nil,"Wrong brand"
}