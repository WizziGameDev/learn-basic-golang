package main

import "fmt"

func main() {

	var testArray [3]string
	testArray[0] = "Hello"
	testArray[1] = "World"
	testArray[2] = "Boy"

	fmt.Println(testArray[1])

	// Define simple array (can't to append)
	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
