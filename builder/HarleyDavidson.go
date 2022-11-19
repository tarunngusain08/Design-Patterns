package builder

type HarleyDavidson struct {
	category string
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

func (b *HarleyDavidson) setCategory(category string) {
	b.category = category
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
		category: b.category,
		suspension: b.suspension,
		engine: b.engine,
		price: b.price
	}
}
