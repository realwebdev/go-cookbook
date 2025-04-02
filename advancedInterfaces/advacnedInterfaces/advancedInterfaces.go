package main

// 1, advanced interface mechanics
// 2, typed functions

// injecting function smartly

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
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

func updateValue(id int, val any, store Storage) error {
	// store := &FooStorage{}
	return store.Put(id, val)
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	// instantiate the server

	updateValue(1, "foo", s.store)
}
