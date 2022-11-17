package main

type ISuitFactory interface {
	makeShirt()
	makePants()
	makeBlazer()
}