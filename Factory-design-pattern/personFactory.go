package	factory

func getPerson(personType string)(IPerson,error){
	if personType == "Customer"{
		return newCustomer(),nil
	}
	if personType == "Employee"{
		return newEmployee(),nil
	}
	return nil,"Wrong person type"
}