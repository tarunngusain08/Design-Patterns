package builder

type Ducati struct {
	category string
	suspension string
	engine string
	price int
}

// func newDucati(suspension string, engine string, price int) *HarleyDavidson {
// 	return &Ducati{
// 		suspension: suspension, 
// 		engine: engine, 
// 		price: price
// 	}
// }

func newDucati() *HarleyDavidson {
	return &Ducati{}
}

func (b *Ducati) setCategory(category string) {
	b.category = category
}

func (b *Ducati) setSuspension(suspension string) {
	b.suspension = suspension
}

func (b *Ducati) setEngine(engine string) {
	b.engine = engine
}

func (b *Ducati) setPrice(price int) {
	b.price = price
}

func (b *Ducati) getBike() *Bike {
	return &Bike{
		category: b.category,
		suspension: b.suspension,
		engine: b.engine,
		price: b.price
	}
}