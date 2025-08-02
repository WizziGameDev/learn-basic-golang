package main

import "fmt"

func main() {

	// Alias untuk int64 dengan nama KTPINT
	type KTPINT int64
	var myKTP KTPINT = 8398299

	fmt.Println(myKTP)
}
