package main

import "fmt"

// func for flexible input value (must be same type data)
func variadicFunc(number ...int) int {
	total := 0

	for _, v := range number {
		total += v
	}

	return total
}

func main() {
	count := variadicFunc(3, 4, 5)
	fmt.Println(count)

	// If we have slice we can add in parameter using nameVariable...
	testSum := []int{1, 2, 3, 4, 5}
	testCount := variadicFunc(testSum...)
	fmt.Println(testCount)
}
