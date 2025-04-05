package main

import "fmt"

func reverseSlice(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println("before", a)
	b := reverseSlice(a)
	fmt.Println("after", b)
}
