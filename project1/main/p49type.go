package main

import "fmt"

/**
* 类型别名
**/
func main() {
	//type MyInt int
	//var i MyInt
	//i = 100
	//// i: main.MyInt, 100
	//fmt.Printf("i: %T, %v\n", i, i)

	type MyInt = int
	var i MyInt
	i = 100
	// i: int, 100
	fmt.Printf("i: %T, %v\n", i, i)

}
