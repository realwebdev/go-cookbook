package main

import "fmt"

// swap adjacent element

func main() {
	// adjacent should remain with in the bounds
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(a)-1; i += 2 {
		a[i], a[i+1] = a[i+1], a[i]
	}

	fmt.Println("adjacent swapping", a)

}
