package main

import (
	"fmt"
	"runtime"
	"time"
)

// gosched让出当前时间片
//func gosched() {
//	go goschedeg("java")
//	for i := 0; i < 2; i++ {
//		runtime.Gosched()
//		fmt.Println("golang")
//	}
//}

//func goschedEg(msg string) {
//	for i := 0; i < 2; i++ {
//		fmt.Printf("msg %v %v\n", msg, i)
//	}
//}

func goExited() {
	go goExitedEg()
	time.Sleep(time.Second)
}

func goExitedEg() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func a(str string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v:%v\n", str, i)
	}
}
func b(str string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v:%v\n", str, i)
	}
}

func main() {
	//gosched()
	//goExited()

	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go a("a")
	go b("b")
	time.Sleep(time.Second * 2)
}
