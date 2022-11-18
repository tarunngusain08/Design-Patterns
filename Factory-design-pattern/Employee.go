package factory

type Employee struct{
	Person
}

func newEmployee() IPerson{
	return &Employee{
		Person: Person{
			name: "Tarun Gusain",
			age: 23,
		}
	}
}