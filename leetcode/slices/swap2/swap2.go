package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	// swapping first and the last element of the slice
	a[0], a[len(a)-1] = a[len(a)-1], a[0]
	fmt.Println(a)

}
