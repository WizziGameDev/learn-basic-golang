package main

import "fmt"

// key-value, jika ada key yang sama maka value yang digunakan adalah yang baru
func main() {

	// data type -> map[key]value
	person := map[string]string{
		"name": "John Doe",
		"age":  "25",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	number := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(number["one"])
	fmt.Println(number["two"])
	fmt.Println(number["three"])
	fmt.Println(number)

	// Function pada MAP
	testMap := map[string]string{
		"name":    "Adam Jordan",
		"age":     "21",
		"address": "Surabaya",
	}

	// Print map
	fmt.Println(testMap)

	// Length map
	fmt.Println(len(testMap))

	// Take key
	fmt.Println(testMap["name"])

	// Change value by key
	testMap["name"] = "Budi"
	fmt.Println(testMap)

	// Create new map with another way
	makeMap := make(map[string]int)
	makeMap["dia"] = 1
	makeMap["kamu"] = 2
	fmt.Println(makeMap)

	// Delete map via key
	delete(makeMap, "dia")
	fmt.Println(makeMap)
}
