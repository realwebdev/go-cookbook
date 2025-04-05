package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3, 6},
		{4, 5, 6, 6},
		{7, 8, 9, 6},
	}

	sum := 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			sum += a[i][j]
		}
	}

	fmt.Println("the sum is ", sum)
}
