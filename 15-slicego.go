package main

import "fmt"

func main() {

	var values = []string{
		"Budi", "Joko", "Samantha", "Berdi", "Wahyu", "Ferdi",
	}

	// [1:4] ambil data index ke-1 dan sebelum index ke-4
	slice := values[1:4]

	fmt.Println(slice)
	// result: [Joko Samantha Berdi]

	// Ambil index ke-2 sampai terakhir
	slice2 := values[2:]
	fmt.Println(slice2)

	// Ambil dari awal sampai sebelum index ke-4
	slice3 := values[:4]
	fmt.Println(slice3)

	// Operasi Slice
	Days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := Days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	/*
		Jika sudah habis panjangnya dia akan membuat array baru, jdi ketika days
		diubah maka daySlice2 tidak akan terpengaruh karena itu array baru
	*/

	daySlice2 := append(daySlice1, "Hari Libur")
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(Days)

	// Membuat slice langsung
	newSlice := make([]int, 2, 5) // -> 3 length, 5 capacity

	newSlice[0] = 1
	newSlice[1] = 2
	fmt.Println(newSlice)
	fmt.Println(len(newSlice), cap(newSlice))

	// Masih menggunakna array yang sama karena length tidak melebihi 5/capacity
	newSlice2 := append(newSlice, 1, 2)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2), cap(newSlice2))

	// Copy slice
	fromSlice := Days[:] // ambil semua days nya
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// Membedakan array dan slice pada (...) didalam []
	iniArray := [...]int{1, 2, 3} // array
	iniSlice := []int{1, 2, 3}    // slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
