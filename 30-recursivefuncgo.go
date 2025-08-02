package main

import "fmt"

// Sebuah cara untuk memanggil func dirinya sendiri
func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialRecursive(number-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
