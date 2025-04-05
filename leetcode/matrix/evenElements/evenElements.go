package main

import "fmt"

func main() {

	a := [][]int{
		{1, 2, 3, 6},
		{4, 5, 6, 6},
		{7, 8, 9, 6},
	}

	// print even elements

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j]%2 == 0 {
				fmt.Printf("the even elements are a[%d][%d] --> %d\n", i, j, a[i][j])
			}
		}
	}

}
