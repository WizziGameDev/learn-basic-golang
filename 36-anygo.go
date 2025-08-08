package main

import "fmt"

// any/interface{} untuk melakukan return data bebas
func Name() any {
	return "dimas"
}

func main() {
	fmt.Println(Name())
}
