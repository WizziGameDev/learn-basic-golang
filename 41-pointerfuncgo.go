package main

import "fmt"

type Phone struct {
	name, model string
	seri        int
}

// pointer function
func updateSeri(phone *Phone) {
	phone.seri += 1
}

func main() {
	// efisien ini langsung jadi pointer
	myPhone := &Phone{"Iphone", "IP", 12}
	updateSeri(myPhone) // pointer

	// harus mengambil alamat
	//myPhone := Phone{"Iphone", "IP", 12}
	//updateSeri(&myPhone) // pointer

	fmt.Println(myPhone)
}
