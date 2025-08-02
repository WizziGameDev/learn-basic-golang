package main

import "fmt"

// Deklarasikan variable di return nya
func testHello() (first, second string) {
	first = "Hello"
	second = "World"

	// must be same for variable
	return first, second
}

func main() {
	fmt.Println(testHello())
}
