package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("timer1 1st time.Now() ", time.Now())
	// 阻塞，直到指定时间
	//t1 := <-timer.C
	//fmt.Println("t1", t1)
	<-timer1.C
	fmt.Println("timer1 2nd time.Now() ", time.Now())

	<-time.After(time.Second * 2)
	fmt.Println("time.Now() ", time.Now())

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("timer2 1st time.Now() ", time.Now())
	}()

	timer2.Stop()
	fmt.Println("timer2 2nd time.Now() ", time.Now())

	time3 := time.NewTimer(time.Second * 5)
	time3.Reset(time.Second * 1)
	<-time3.C
	fmt.Println("timer3 1st time.Now() ", time.Now())

}
