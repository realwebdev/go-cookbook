package main

import "fmt"

// 1, advanced interface mechanics
// 2, typed functions
// 3. injecting function smartly

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}

type simplePutter struct {
}

func (s *simplePutter) Put(id int, val any) error {
	return nil
}
func (s *simplePutter) String() string { return "" }

type Storage interface {
	Putter
	Get(id int) (any, error)
}

type FooStorage struct {
}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct {
}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	// store := &FooStorage{}
	return p.Put(id, val)
}

func main() {
	sputter := &simplePutter{}
	// instantiate the server

	updateValue(1, "foo", sputter)
}
