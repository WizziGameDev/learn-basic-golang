package main

import "fmt"

func main() {

	// Switch Case usually use for checking in 1 variable
	name := "Adam"

	switch name {
	case "Adam":
		fmt.Println("I'An Adam")
	case "Joko":
		fmt.Println("I'An Joko")
	default:
		fmt.Println("I'An Idea")
	}

	switch len(name) > 10 {
	case true:
		fmt.Println("I'An 10")
	case false:
		fmt.Println("I'An min 10")
	}
}
