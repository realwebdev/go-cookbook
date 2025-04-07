package main

import "fmt"

/*

Create a hash map for rows to store their frequency.

Transpose the matrix so columns become rows.

Count identical pairs by comparing the rows with the transposed columns.
*/

func equalPairs(grid [][]int) int {
    n := len(grid)
    rowMap := make(map[string]int)

    // Step 1: Store all rows as string keys
    for i := 0; i < n; i++ {
        key := fmt.Sprint(grid[i])
        rowMap[key]++
    }

    count := 0

    // Step 2: For each column, convert it to string and check against rowMap
    for j := 0; j < n; j++ {
        col := []int{}
        for i := 0; i < n; i++ {
            col = append(col, grid[i][j])
        }
        key := fmt.Sprint(col)
        count += rowMap[key] // add how many times this "column" appears as a row
    }

    return count
}


func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	grid2 := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}

	ans1 := equalPairs(grid)
	ans2 := equalPairs(grid2)

	fmt.Printf("ans1 --> %d\n, ans2 --> %d\n", ans1, ans2)

}
