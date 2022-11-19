package builder

type IBikeBuilder interface{
	makeCruiser() ICruiser
	makeSuperBike() ISuperBike
}

func GetBikeBuilder(name string) IBikeBuilder{
	if name == "HarleyDavidson"{
		return &HarleyDavidson{}
	} 
	if name == "Ducati"{
		return &Ducati{}
	}
	return nil
}