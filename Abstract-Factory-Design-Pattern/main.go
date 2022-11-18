package abstract_factory

func main(){
	peterEnglandFactory,_ := GetSuitFactory("PeterEngland")
	levisFactory,_ := GetSuitFactory("Levis")

	PeterEnglandPants := peterEnglandFactory.makePants()
	PeterEnglandShirt := peterEnglandFactory.makeShirt()
	PeterEnglandBlazer := peterEnglandFactory.makeBlazer()

	LevisPants := levisFactory.makePants()
	LevisShirt := levisFactory.makeShirt()
	LevisBlazer := levisFactory.makeBlazer()
	
	printPantsDetails(PeterEnglandPants)
	printShirtDetails(PeterEnglandShirt)
	printBlazerDetails(PeterEnglandBlazer)

	printPantsDetails(LevisPants)
	printShirtDetails(LevisShirt)
	printBlazerDetails(LevisBlazer)
}

func printPantsDetails(pants IPants){
	fmt.Sprintf("%v pants is of %v size",pants.logo,pants.size)
}

func printShirtDetails(shirt IShirt){
	fmt.Sprintf("%v shirt is of %v size",shirt.logo,shirt.size)
}

func printBlazerDetails(blazer IBlazer){
	fmt.Sprintf("%v blazer is of %v size",blazer.logo,blazer.size)
}