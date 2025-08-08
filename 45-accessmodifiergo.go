package main

import (
	"fmt"
	"go-dasar/helper"
)

/*
	var / function jika penamaannya di awali dengan huruf kecil
	hanya dapat diakses di package itu saja,tidak bisa di akses diluar package.

	Kalau memakai huruf besar diawalannya dapat diakses diluar package
*/

func main() {
	//fmt.Println(helper.sayHello()) // LUAR PACKAGE
	fmt.Println(helper.Name) // Ini bisa karena awalan huruf besar
}
