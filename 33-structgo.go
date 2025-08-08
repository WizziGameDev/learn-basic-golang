package main

import "fmt"

// mirip class object di java dan juga seperti alias di GO
type Customer struct {
	Firstname, Lastname string
	Age                 int
}

func main() {
	var dataCustomer Customer
	dataCustomer.Firstname = "Gopher"
	dataCustomer.Lastname = "Ino"
	dataCustomer.Age = 25

	fmt.Println(dataCustomer.Firstname)
	fmt.Println(dataCustomer.Lastname)
	fmt.Println(dataCustomer.Age)

	// Struct Literal
	budi := Customer{
		Firstname: "Budi",
		Lastname:  "Anduk",
		Age:       40,
	}
	fmt.Println(budi.Firstname)
	fmt.Println(budi.Lastname)
	fmt.Println(budi.Age)

	joko := Customer{"Joko", "Dima", 30}
	fmt.Println(joko.Firstname)
	fmt.Println(joko.Lastname)
	fmt.Println(joko.Age)
}
