package main

// defer will be run after function done
func deferAlert() string {
	return "CODE DONE"
}

func testDefer() string {
	defer deferAlert()
	print("Test defer alert")
	return "Defer after this print"
}

// panic will be run and code really break/stop
func panicAlert() string {
	return "ERROR CODE"
}

func testPanic(test bool) string {
	if test {
		return panicAlert()
	}
	return "Test panic"
}

func main() {
	testDefer()
	testPanic(true)
}
