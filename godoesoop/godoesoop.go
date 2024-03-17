package main

import "fmt"

// abstraction
// encapsulation
// inheritance
// polymorphism

type NumberStorer interface {
	// desrcibe what interface can do?
	// football player can football / kick a ball / knows the rules
	GetAll() ([]int, error)
	Put(int) error
} // to implement this interface, you need to implement these methods

type MongoDbNumberStore struct {
	// mongodb client
	// some values
}

type PostgresNumberStore struct {
	// postgres client
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil // nil is defined as nothing
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the postgres stoage") // nil is defined as nothing
	return nil
}

func (m MongoDbNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil // nil is defined as nothing
}

func (m MongoDbNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB stoage") // nil is defined as nothing
	return nil
}

type ApiServer struct {
	numberStore NumberStorer // a losely coupled system
	// numberStore MongoStore (a bad behvaior)
}

func main() {
	apiServer := ApiServer{
		numberStore: PostgresNumberStore{},
	}

	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)

}
