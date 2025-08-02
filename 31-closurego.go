package main

import "fmt"

// Mudah untuk dibuat namun sangat tidak disarankan jika code sudah besar agar tidak membingungkan

func main() {
	count := 0

	incrementClosure := func() {
		fmt.Println("incrementClosure")
		count++
	}

	incrementClosure()
	incrementClosure()
	fmt.Println(count)
}
