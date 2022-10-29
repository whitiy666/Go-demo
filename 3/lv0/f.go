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
		channel <- 4
		fmt.Println("发送第四个数据，但是肯定发不出去，已经超过缓存空间3")
	}()

	time.Sleep(time.Second)
}
