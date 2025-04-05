// a := [][]int{
// 	{1, 2, 3, 6},
// 	{4, 5, 6, 6},
// 	{7, 8, 9, 6},
// }

package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3, 6},
		{4, 5, 6, 6},
		{7, 8, 9, 6},
		{7, 8, 9, 6},
	}

	if len(a) == len(a[0]) {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[i]); j++ {
				if i == j {
					fmt.Printf("the diagnol of matrix are a[%d][%d] -->%d\n", i, j, a[i][j])
				}

			}
		}

	}
}
