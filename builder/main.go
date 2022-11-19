package builder

func main(){
	category := "Cruiser"
	cruiserBuilder := GetBikeBuilder(category")
	cruiserBuilder.setSuspension("Cruiser Suspension")
	cruiserBuilder.setEngine("Cruiser Engine")
	cruiserBuilder.setPrice(1500000)
	cruiserBuilder.setCategory(category)
	printBikeDetails(cruiserBuilder.getBike())


	category = "SportsBike"
	sportsBikeBuilder := GetBikeBuilder(category)
	sportsBikeBuilder.setSuspension("SportsBike Suspension")
	sportsBikeBuilder.setEngine("SportsBike Engine")
	sportsBikeBuilder.setPrice(3500000)
	sportsBikeBuilder.setCategory(category)
	printBikeDetails(sportsBikeBuilder.getBike())
}

func printBikeDetails(category string, bike Bike) {
	fmt.Sprintf("The %v bike of %v category is ready with brand 
	new %v suspension and %v engine with an amazing price of
	 %v", bike.brand, bike.category, bike.suspension, bike.engine, bike.price)
}