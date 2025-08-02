package main

import "fmt"

func main() {

	count := 0

	for count < 10 {
		fmt.Println(count)
		count++
	}

	// For Statement
	for x := 0; x < 10; x++ {
		fmt.Println("Test ke-", x)
	}

	// Iteration pada array, slice or map
	data := []int{1, 2, 3, 4, 5}

	// manual
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	// Use range for ForEach
	// if u use index, u need to declare
	for index, value := range data {
		fmt.Println(index, " : ", value)
	}

	// If u not use index u need to declare (_)
	for _, value := range data {
		fmt.Println(value)
	}
}
