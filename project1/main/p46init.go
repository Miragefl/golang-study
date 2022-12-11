package main

import "fmt"

var i int = initVar()

func initVar() int {
	fmt.Println("initVar ...")
	return 100
}

func init() {
	fmt.Println("init ...")
}
func init() {
	fmt.Println("init1 ...")
}

func main() {
	fmt.Println("main ...")
}
