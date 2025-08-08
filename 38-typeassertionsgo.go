package main

import (
	"fmt"
	"reflect"
)

// Melakukan konversi dari return function
func data() any {
	return "hello world"
}

func main() {
	dataLama := data()
	dataString := dataLama.(string)
	fmt.Println(dataString)

	// Conversi salah jika type yang di konversi tidak sesuai
	dataSalah := dataLama.(int) // panic
	fmt.Println(dataSalah)

	// Jika panic code akan berhenti, bisa bakai switch case
	switch dataSwitch := dataLama.(type) {
	case int:
		fmt.Println("Data int", dataSwitch)
	case string:
		fmt.Println("Data string", dataSwitch)
	default:
		fmt.Println("Unknow data type", reflect.TypeOf(dataLama))
	}
}
