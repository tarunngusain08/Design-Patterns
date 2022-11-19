package builder

type IBikeBuilder interface{
	setSuspension(suspension string)
	setEngine(engine string)
	setPrice(price int)
	getBike() *Bike
}

func GetBikeBuilder(category string) IBikeBuilder{
	if category == "Cruiser"{
		return newHarleyDavidson()
	} 
	if category == "SportsBike"{
		return newDucati()
	}
	return nil
}