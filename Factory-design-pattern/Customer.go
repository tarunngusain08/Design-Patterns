package	factory

type Customer struct {
	Person
}

func newCustomer() IPerson{
	return &Customer{
		Person: Person{
			name: "Tarun",
			age: 23
	}
}