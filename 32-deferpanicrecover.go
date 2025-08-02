package main

import "fmt"

// defer will be run after function done
func deferAlert() {
	fmt.Println("Code End")
}

// panic will be run and code really break/stop
func runApp(test bool) string {
	defer deferAlert()
	if test {
		panic("Panic code break!")
	}
	return "Test panic"
}

func main() {
	runApp(true)
}
