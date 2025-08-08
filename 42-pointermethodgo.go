package main

import "fmt"

type PointMethod struct {
	name string
}

func (p *PointMethod) hello(name string) {
	p.name = "hello " + name
}

func main() {
	data := &PointMethod{}
	data.hello("world")
	fmt.Println(data.name)
}
