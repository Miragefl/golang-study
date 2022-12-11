package main

import "fmt"

func f1() {
	fmt.Println("start...")
	fmt.Println("step1...")
	fmt.Println("end...")
}
func testDefer() {
	fmt.Println("start")
	defer fmt.Println("step1 ....")
	defer fmt.Println("step2 ....")
	defer fmt.Println("step3 ....")
	fmt.Println("end")
}
func main() {
	f1()
	fmt.Println("------------")
	testDefer()
}
