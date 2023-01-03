package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	//counter := 1
	//for _ = range ticker.C {
	//	fmt.Println("ticker time.Now() ", time.Now())
	//	counter++
	//	if counter >= 5 {
	//		break
	//	}
	//}
	//ticker.Stop()
	chanInt := make(chan int)

	go func() {
		for range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}

	}()

	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Printf("接受到的参数: %v ,当前sum值: %v\n", v, sum)
		if sum > 10 {
			break
		}
	}

}
