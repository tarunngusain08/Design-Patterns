package builder

type HarleyDavidson struct {
	suspension string
	engine string
	price int
}

// func newHarleyDavidson(suspension string, engine string, price int) *HarleyDavidson {
// 	return &HarleyDavidson{
// 		suspension: suspension, 
// 		engine: engine, 
// 		price: price
// 	}
// }
func newHarleyDavidson() *HarleyDavidson {
	return &HarleyDavidson{}
}

func (b *HarleyDavidson) setSuspension(suspension string) {
	b.suspension = suspension
}

func (b *HarleyDavidson) setEngine(engine string) {
	b.engine = engine
}

func (b *HarleyDavidson) setPrice(price int) {
	b.price = price
}

func (b *HarleyDavidson) getBike() *Bike {
	return &Bike{
		suspension: b.suspension,
		engine: b.engine,
	price: b.price
	}
}
