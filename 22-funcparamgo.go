package main

import "fmt"

func sayHelloBro(firstname string, lastname string) {
	fmt.Println("Hello " + firstname + " " + lastname)
}

func main() {
	sayHelloBro("Adam", "Saputra")
}
