package main

import "fmt"

type Detail struct {
	firstName, lastName string
}

/*
	logic dari asterisk adalah ketika ada pointer(reference),
	reference variable yang di beri asterisk(*) akan melakukan
	reference ke object baru yang dibuat oleh variable tersebut
*/

func main() {
	data1 := Detail{"Budi", "Anduk"}
	data2 := &data1

	data2.lastName = "Utomo"

	fmt.Println(data1)
	fmt.Println(data2)

	// asterisk operator (*)
	*data2 = Detail{"Dimas", "Mabar"}
	fmt.Println(data1)
	fmt.Println(data2)
}
