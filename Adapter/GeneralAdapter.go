package main

import "fmt"

// Client code
func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)
	fmt.Println(adapter.Request())
}

// Adaptee is the legacy interface
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee method called"
}

// Target is the new interface
type Target interface {
	Request() string
}

// Adapter is the concrete adapter
type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
	return fmt.Sprintf("Adapter method called with %s", a.adaptee.SpecificRequest())
}
