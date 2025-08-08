package main

import "fmt"

type User struct {
	firstName, lastName string
}

func main() {
	data1 := User{"Dimas", "Joko"}
	data2 := data1 // copy
	// data 1 tidak akan berubah karena masih type pass by value
	data2.firstName = "Dilan"
	fmt.Println(data1)
	fmt.Println(data2)

	// Pointer
	user1 := User{"Bang", "Malaka"}
	user2 := &user1 // pointer

	// pass by reference, yang mana object head, ketika ada yang mengubah di turunannya/function lain
	// akan ikut berubah di object head nya (kayak di java)
	user2.firstName = "Tan"
	fmt.Println(user1)
	fmt.Println(user2)
}
