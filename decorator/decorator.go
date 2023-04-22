package main

import (
	"fmt"
)

// go get gopkg.in/couchbase/gocb.v1
func Decorate(f func(string)) func(string) {
	return func(name string) {
		fmt.Println("Before function call")
		f(name)
		fmt.Println("After function call")
	}
}

func Hello(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	decoratedHello := Decorate(Hello)
	decoratedHello("John")
}

/*
In this example, DoubleOutput is a higher-order function that takes a function f as input and 
returns a new function that doubles the output of f. We use this higher-order function to create 
a new function doubleSquare, which is a decorated version of the Square function that doubles 
its output. We then call doubleSquare with an input of 3, which outputs 18.

Note that the input function f must take an integer input and return an integer output, which 
is enforced by the function signature of DoubleOutput. This is necessary for the returned function 
to be able to operate on f and return the doubled output.
*/
