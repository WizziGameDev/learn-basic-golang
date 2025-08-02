package main

import "fmt"

func main() {
	// Perlu diperhatikan untuk digit value jangan melebihi dari konversi di bawahnya
	var nilai32 int32 = 4554
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	// Konversi dari String ke uint -> String Lagi
	var name string = "Jordan"
	var J uint8 = name[0]
	var JString string = string(J)

	fmt.Println(name)
	fmt.Println(J)
	fmt.Println(JString)
}
