package main

import (
	"fmt"
)

/*
	Cara kerja nya adalah untuk func yang memiliki parameter yang berisikan func harus memiliki input dan return yang
	sama saat membuat/memasukkan ke parameter. Jika kita membuat func dengan input string dan return string, dan
	ingin memasukkan ke func parameter kita perlu samain dengan kondisi input dan return nya
*/

func sayHelloWithFilter(firstName string, filter func(string) string) string {
	filteredName := filter(firstName)
	return "Success filter with name " + filteredName
}

func spamFilter(firstName string) string {
	if firstName == "Anjing" {
		return "..."
	} else {
		return firstName
	}
}

func main() {
	myName := "Adam"

	testFilter := sayHelloWithFilter(myName, spamFilter)
	fmt.Println(testFilter)
}
