package main

import "fmt"

// nill digunakan hanya untuk beberapa tipe data
// interface/any, function, map, slice, pointer, dan channel

func dataNil(name string) interface{} {
	if name == "" {
		return nil
	} else {
		return name
	}
}

func main() {
	fmt.Println("Hello")
}
