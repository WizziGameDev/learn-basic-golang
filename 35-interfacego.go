package main

import "fmt"

// Kontrak tanpa perlu secara explisit melakukan implements seperti di java
type testInterface interface {
	SayHello() string
	SayGoodbye() string
}

func main() {
	data := Person{Name: "adam"}
	fmt.Println(data.SayHello())
}

type Person struct {
	Name string
}

func (p Person) SayHello() string {
	return "Hello " + p.Name
}

func (p Person) SayGoodbye() string {
	return "Goodbye " + p.Name
}
