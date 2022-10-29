package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //无缓存通道，利用阻塞特性。没有接受就阻塞
	go gA(ch)
	go gB(ch)

	time.Sleep(time.Second)
}

func gA(p chan int) {
	for i := 1; i <= 100; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Printf("A线程：%d\n", i) //奇数
		}
	}
}

func gB(p chan int) {
	for i := 1; i <= 100; i++ {
		<-p
		if i%2 == 0 {
			fmt.Printf("B线程：%d\n", i) //偶数
		}
	}
}
