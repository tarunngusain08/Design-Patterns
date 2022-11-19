package builder

type IBikeBuilder interface{
	setSuspension(suspension string)
	setEngine(engine string)
	setPrice(price int)
	getBike() *Bike
}

func GetBikeBuilder(name string) IBikeBuilder{
	if name == "HarleyDavidson"{
		return newHarleyDavidson()
	} 
	if name == "Ducati"{
		return newDucati()
	}
	return nil
}