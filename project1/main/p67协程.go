package main

import (
	"fmt"
	"time"
)

func main() {
	go show("java")
	go show("golang")
	time.Sleep(time.Second * 5)
}

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v \n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
