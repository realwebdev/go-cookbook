package main

import "fmt"

func equalPairs(matrix [][]int) int {
	// step 1: create a map to count the frequency of rows
	rowMap := make(map[string]int)

	// step2: convert each row to a string and store in the map
	for z, row := range matrix {
		fmt.Println(z)
		// convert the row to a string or tuple
		rowStr := fmt.Sprintf("%v", row)
		rowMap[rowStr]++
	}

	// step 3: Transpose the matrix and comapare with the rows
	count := 0

	n := len(matrix)
	for j := 0; j < n; j++ {
		// form a coloumn as a string
		coloumn := make([]int, n)
		for i := 0; i < n; i++ {
			coloumn[i] = matrix[i][j]
		}

		// convert the coloumn to a string
		coloumnStr := fmt.Sprintf("%v", coloumn)
		count += rowMap[coloumnStr]
	}

	// step 4: Return the total count of equal pairs
	return count
}

func main() {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(equalPairs(matrix))

}
