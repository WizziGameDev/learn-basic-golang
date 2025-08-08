package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi is zero")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, err := pembagian(1, 0)
	if err == nil {
		fmt.Println("Pembagian berhasil, Hasil: ", hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
