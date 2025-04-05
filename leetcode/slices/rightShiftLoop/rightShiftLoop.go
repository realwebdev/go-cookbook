package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	last := a[len(a)-1]

	// shifting all elements to right

	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i-1]
	}

	a[0] = last

	fmt.Println(a)
}
