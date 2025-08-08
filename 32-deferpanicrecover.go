package main

import "fmt"

// defer will be run after function done
func deferAlert() {
	fmt.Println("Code End")

	// catch and recovery
	message := recover()
	fmt.Println(message)
}

// panic will be run and code really break/stop
func runApp(test bool) string {
	defer deferAlert()
	if test {
		panic("panic: Panic code break!")
	}
	return "Test panic"
}

// Dengan recover code akan berjalan terus dan kita akan tahu panic nya karena apa
func main() {
	runApp(true)
}
