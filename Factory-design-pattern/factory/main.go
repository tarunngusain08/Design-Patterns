package main

func main(){
	employee,_ := getPerson("employee")
	customer,_ := getPerson("customer")
}

func printDetails(p IPerson) {
	fmt.Printf("Person %v\n",p.getName())
	fmt.Printf("Age %v\n",p.getAge())
}