// package main

// import "fmt"

// type NumberStorer interface {
// 	GetAll() ([]int, error)
// 	Put(int) error
// } // let just integrate postgres example with this

// type PostgresNumberStore struct {
// 	// some values
// } // it will recieve the methods of the interface

// func (p PostgresNumberStore) Put(int) error {
// 	fmt.Println("adding the number to the postgres database")
// 	return nil
// }

// func (p PostgresNumberStore) GetAll() ([]int, error) {
// 	return []int{1, 2, 3}, nil
// }

// // to implement an interface, you need to implement all the methods of the interface
// // if you don't implement all the methods, you will get an error

// type MongoDBNumberStore struct {
// 	// some values
// } // it will recieve the methods of the interface

// func (m MongoDBNumberStore) Put(int) error {
// 	fmt.Println("adding the number to the database")
// 	return nil
// }

// func (m MongoDBNumberStore) GetAll() ([]int, error) {
// 	return []int{1, 2, 3}, nil
// }

// type ApiServer struct {
// 	NumberStorer NumberStorer
// }

// func main() {
// 	apiServer := ApiServer{
// 		NumberStorer: PostgresNumberStore{}, // constructor
// 	} // this will work because MongoDBNumberStore implements the NumberStorer interface. it is a polymorphism concept
// 	apiServer.NumberStorer.Put(1)
// 	number, err := apiServer.NumberStorer.GetAll()
// 	if err != nil {
// 		panic(err)

// 	}

// 	fmt.Println(number)
// }

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array[2:])
}
