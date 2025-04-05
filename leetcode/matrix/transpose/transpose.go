// swapping rows with coloumsn
package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)

	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}

func main() {
	a := [][]int{{1, 2, 3, 6}, {4, 5, 6, 6}, {7, 8, 9, 6}}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		for j := 0; j < len(a); j++ {
			fmt.Println(a[j])
		}
	}

	d := transpose(a)

	fmt.Println(d)

}
