package prototype

import "fmt"

// Prototype interface
type Person interface {
    Clone() Person
}

// Concrete prototype struct
type Employee struct {
    Name    string
    Age     int
    Address string
}

// Clone method creates a new copy of the Employee struct
func (e *Employee) Clone() Person {
    return &Employee{
        Name:    e.Name,
        Age:     e.Age,
        Address: e.Address,
    }
}

func main() {
    // Create an Employee object
    emp1 := &Employee{
        Name:    "John Doe",
        Age:     30,
        Address: "123 Main St",
    }

    // Clone the Employee object
    emp2 := emp1.Clone().(*Employee)

    // Modify the cloned Employee object
    emp2.Name = "Jane Smith"
    emp2.Age = 25
    emp2.Address = "456 Elm St"

    // Print the original and cloned Employee objects
    fmt.Println("Original Employee:", emp1)
    fmt.Println("Cloned Employee:", emp2)
}
