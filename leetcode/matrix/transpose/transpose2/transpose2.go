package main

import "fmt"

func main() {

	a := [][]int{
		{1, 2, 3, 6},
		{4, 5, 6, 6},
		{7, 8, 9, 6},
		{7, 8, 9, 6},
	}

	rows := len(a)
	cols := len(a[0])

	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	// fill with swapped positions
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = a[i][j]
		}
	}

	for g, row := range result {

		fmt.Println(row, g)
	}
	fmt.Println(result)

}
