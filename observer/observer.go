package main

import (
	"fmt"
	"time"
)

type Observer interface {
	Notify(string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) UnregisterObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		go observer.Notify(message)
	}
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Notify(message string) {
	fmt.Printf("%s received message: %s\n", o.name, message)
}

func main() {
	subject := &Subject{}
	observer1 := &ConcreteObserver{"Observer 1"}
	observer2 := &ConcreteObserver{"Observer 2"}

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.NotifyObservers("Hello, observers!")
	time.Sleep(time.Second)

	subject.UnregisterObserver(observer1)

	subject.NotifyObservers("Goodbye, observer 1!")
	time.Sleep(time.Second)
}
