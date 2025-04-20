package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)

	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 { // check if the current plot is 0
			leftEmpty := (i == 0) || (flowerbed[i-1] == 0) // check if the left neighbor is empty or if it's the first plot no left neighbor
			rightEmpty := (i == length-1) || (flowerbed[i+1] == 0) // check if the right neighbot is empty or if its the last plot no right neighbor

			if leftEmpty && rightEmpty {
				flowerbed[i] = 1 // plant the flower
				count++
				if count >= n {
					return true
				}
			}
		}
	}

	return count >= n

}

func main() {

	flowerBed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerBed, n))

}
