package main

import "fmt"

func main() {

	name := "Adam"

	if name == "Adam" {
		fmt.Println("I'Am Adam")
	} else if name != "Joko" {
		fmt.Println("I'Am Joko")
	} else {
		fmt.Println("I'An Idea")
	}

	// Short if statement
	if length := len(name); length > 0 {
		fmt.Println("Length name = ", length)
	} else {
		fmt.Println("Length = 0")
	}
}
