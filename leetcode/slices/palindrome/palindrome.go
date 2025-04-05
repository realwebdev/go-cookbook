package main

import "fmt"

func main() {

	// length of array should be even

	// palindrom comparing inwards
	a := []int{1, 2, 3, 2, 1, 4}

	isPal := true

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		if a[i] != a[j] {
			isPal = false
			break
		}
	}

	fmt.Println("a is -->", isPal)

}
