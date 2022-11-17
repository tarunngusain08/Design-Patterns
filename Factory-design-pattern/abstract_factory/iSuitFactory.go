package abstract_factory

type ISuitFactory interface {
	makeShirt() IShirt
	makePants() IPants
	makeBlazer() IBlazer
}