package main

import "fmt"

func main() {
	mobile := Mobile{
		name: "m",
	}
	mobile.read()
	mobile.write()

	computer := Computer{
		model: "dell",
	}
	computer.read()
	//computer.write()
}

type USB interface {
	read()
	write()
}

type Computer struct {
	model string
}
type Mobile struct {
	name string
}

func (mobile Mobile) read() {
	fmt.Println("mobile read ....")
}

func (mobile Mobile) write() {
	fmt.Println("mobile write ....")
}

func (c Computer) read() {
	fmt.Printf("model: %v reading... \n", c.model)
}

//
//func (c Computer) write() {
//	fmt.Printf("model: %v writing... \n", c.model)
//}
