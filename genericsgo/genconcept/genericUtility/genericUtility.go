package main

import "fmt"

func Contains[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3}, 2))           // true
	fmt.Println(Contains([]string{"go", "dev"}, "go")) // true

}
