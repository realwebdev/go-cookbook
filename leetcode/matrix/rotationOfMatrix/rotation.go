package main

import "fmt"

func main() {
	// rotation of matrix

	a := [][]int{
		{1, 2, 3, 6},
		{4, 5, 6, 6},
		{7, 8, 9, 6},
		{7, 8, 9, 6},
	}

	row := len(a)
	col := len(a[0])

	result := make([][]int, col)
	for i := range result {
		result[i] = make([]int, row)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			result[j][i] = a[i][j]
		}
	}

	// reverse each row of the transpose matrix
	for i := 0; i < len(result); i++ {
		for j, k := 0, len(result[i])-1; j < k; j, k = j+1, k-1 {
			result[i][j], result[i][k] = result[i][k], result[i][j]
		}
	}

	// Print the rotated matrix
	fmt.Println("Rotated Matrix (90 degrees clockwise):")
	for _, row := range result {
		fmt.Println(row)
	}

}
