package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func main() {
	defer close(channel)
	go send()
	fmt.Println("waiting ....")
	go receive()
	time.Sleep(time.Millisecond * 500)
}

func send() {
	//rand.Seed(time.Now().UnixNano())
	//value := rand.Intn(10)
	value := 10
	fmt.Printf("send %v \n", value)
	time.Sleep(time.Millisecond * 100)
	channel <- value
}

func receive() {
	value := <-channel
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end ....")
}
