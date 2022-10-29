package main

import (
	"fmt"
	"time"
)

//带缓存的channel

func main() {
	var channel = make(chan int, 3) //3个空间的channel

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		fmt.Printf("发送了三个数据")
		rec := <-channel //接受一个数据
		fmt.Println(rec)
		channel <- 4
		fmt.Println("发送第四个数据，这次能发送了。没有超过缓存空间3")
	}()

	time.Sleep(time.Second)
}
