package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// left shift

	first := a[0]
	leftShift := append(a[1:], first)

	fmt.Println(leftShift)
}
