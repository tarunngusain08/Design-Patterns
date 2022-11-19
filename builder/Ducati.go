package builder

type Ducati struct {
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

func (b *Ducati) setSuspension(suspension string) {
	b.suspension = suspension
}

func (b *Ducati) setEngine(engine string) {
	b.engine = engine
}

func (b *Ducati) setPrice(price int) {
	b.price = price
}

func (b *HarleyDavidson) getBike() *Bike {
	return &Bike{
		suspension: b.suspension,
		engine: b.engine,
		price: b.price
	}
}