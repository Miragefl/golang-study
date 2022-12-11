package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world")
	testForeach()
}

func testForeach() {
	var s int = 1
	for i := 1; i < 5; i++ {
		s *= i
	}
	fmt.Print(s)
}
