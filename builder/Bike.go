package builder

type IBikeBuilder interface {
	setBrand(logo string) 
	setType(type string)
	setPrice(price int)
	getBrand() string
	getType() string
	getPrice() int
}

type Bike struct {
	brand string
	type string
	price int
}

func (b *Bike) setBrand(brand string) {
	b.brand = brand
}

func (b *Bike) setType(type string) {
	b.type = type
}

func (b *Bike) setPrice(price int) {
	b.price = price
}

func (b *Bike) getBrand()string {
	return b.brand
}

func (b *Bike) getCategory()string {
	return b.category
}

func (b *Bike) getPrice()int {
	return b.price
}
