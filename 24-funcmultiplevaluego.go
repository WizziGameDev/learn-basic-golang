package main

import "fmt"

// Return multi value
func multiValue() (string, string) {
	return "APA", "KAMU"
}

func main() {
	one, two := multiValue()
	fmt.Println(one, two)

	// if you no need lastname, you can define variable with (_)
	first, _ := multiValue()
	fmt.Println(first)
}
