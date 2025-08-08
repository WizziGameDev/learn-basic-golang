package main

import "fmt"

type Data struct {
	Name string
	Time int64
}

func (data Data) SayHello() string {
	return "Hello bro " + data.Name
}

func main() {
	budi := Data{"Budi", 10}
	fmt.Println(budi.SayHello())
}
