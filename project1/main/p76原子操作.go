package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// var sum = 0
//
// var lock sync.Mutex
//
//	func add() {
//		lock.Lock()
//		sum++
//		lock.Unlock()
//	}
//
//	func sub() {
//		lock.Lock()
//		sum--
//		lock.Unlock()
//	}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("sum: %v", i2)
}

var i2 int32 = 100

func add() {
	atomic.AddInt32(&i2, 1)
}
func sub() {
	atomic.AddInt32(&i2, -1)
}
