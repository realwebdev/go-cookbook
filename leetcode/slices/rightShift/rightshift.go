package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// right shift

	last := a[len(a)-1]
	righttShift := append([]int{last}, a[:len(a)-1]...)

	fmt.Println(righttShift)

	/*
	You missed the ... after a[:len(a)-1], which is necessary when appending 
	a slice to another slice in Go.
	Without the ..., you're trying to append a slice as one element, not unpack it.
	*/
}
