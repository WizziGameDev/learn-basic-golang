package main

import "fmt"

func main() {
	var name string

	name = "Adam Jordan"

	fmt.Println("Nama saya: ", name)

	// Langsung
	var age int8 = 20   // define type
	var gender = "male" // no define type

	fmt.Println("Age saya: ", age)
	fmt.Println("Gender saya: ", gender)

	// Membuat variable tanpa harus menggunakan var, deklarasi awal saja
	city := "Surabaya"
	fmt.Println("City saya: ", city)

	// city := "Jombang" tidak bisa, hanya deklarasi awal untuk :=

	// Multiple Variable
	var (
		firstname = "adam"
		lastname  = "jordan"
	)

	fmt.Println(firstname + " " + lastname)
}
