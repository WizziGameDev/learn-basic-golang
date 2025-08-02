package main

import "fmt"

func main() {

	// slice
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// break (stop if condition true)
	for _, value := range data {
		if value == 5 {
			break
		}

		fmt.Println(value)
	}

	// continue (skip if condition true)
	for _, value := range data {
		if value == 5 || value == 9 {
			continue
		}

		fmt.Println(value)
	}
}
