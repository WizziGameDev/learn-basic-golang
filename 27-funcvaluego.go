package main

import "fmt"

func helloApaKabar(firstName string) string {
	return "Hello Apa Kabar " + firstName
}

func main() {
	// To mapping variable to func, no need ()
	sayKabar := helloApaKabar
	fmt.Println(sayKabar("Adam"))
}
