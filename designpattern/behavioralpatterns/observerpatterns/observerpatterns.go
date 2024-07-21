package main

import "fmt"

type Observer interface {
	Update(data string)
}

// define the subject interface and concrete implementation

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(data string)
}

type ConcreteSubject struct {
	observers []Observer
	state     string
}

// register method adds an observer to the list of observers
func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// unregister method removes an observer from the list of observers
func (s *ConcreteSubject) Unregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// notify method sends the data to all the observers
func (s *ConcreteSubject) NotifyAll() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

// get the state of the subject
func (s *ConcreteSubject) GetState() string {
	return s.state
}

// SetState sets the state of the subject and notifies observers
func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.NotifyAll()
}

// Implement the concrete observer
type ConcreteObserverA struct {
	name string
}

// Update method for ConcreteObserverA
func (o *ConcreteObserverA) Update(data string) {
	fmt.Printf("%s received update: %s\n", o.name, data)
}

// ConcreteObserverB struct
type ConcreteObserverB struct {
	name string
}

// Update method for ConcreteObserverB
func (o *ConcreteObserverB) Update(data string) {
	fmt.Printf("%s received update: %s\n", o.name, data)
}

func main() {
	// Create subject
	subject := &ConcreteSubject{}

	// Create observers
	observerA := &ConcreteObserverA{name: "Observer A"}
	observerB := &ConcreteObserverB{name: "Observer B"}

	// Register observers
	subject.Register(observerA)
	subject.Register(observerB)

	// Change state and notify observers
	subject.SetState("State 1")
	subject.SetState("State 2")

	// Deregister an observer
	subject.Unregister(observerA)

	// Change state and notify remaining observers
	subject.SetState("State 3")
}
