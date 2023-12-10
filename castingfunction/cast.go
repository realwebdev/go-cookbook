package main

import "fmt"

type Person struct {
	Name string
}

type Hamza struct {
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

func (h Hamza) SayGoodbye() {
	// This method tries to assert if h is of type Person.
	// If it is, it will call the SayHello() method.
	if person, ok := interface{}(h).(Person); ok {
		person.SayHello()
	} else {
		fmt.Println("I am not a Person.")
	}

	if person1, ok := interface{}(h).(Person); ok {
		person1.SayHello()
	} else {
		fmt.Println("I am not a Person.")
	}
}

func main() {
	p := Person{Name: "John Doe"}

	// Call the SayHello() method on p.
	p.SayHello()

	// Call the SayGoodbye() method on p.
	h := Hamza{}
	h.SayGoodbye()
}
