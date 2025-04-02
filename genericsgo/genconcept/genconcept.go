package main

import "fmt"

// generic can be over kill
// generic data structure

/*
func FunctionName[T TypeConstraint](param T) T { }
*/

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)

	item := s.items[n-1]
	s.items = s.items[:n-1]

	return item
}

func main() {
	// can add stack of any type
	intStack := Stack[int]{}
	intStack.Push(10)

	fmt.Println(intStack.Pop())
	stringStack := Stack[string]{}
	stringStack.Push("hello")
	fmt.Println(stringStack.Pop()) // "hello"

}
