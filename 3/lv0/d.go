package main

//表演一个没有缓存和接收者来使用channel发送数据堵塞的现象

import (
	"fmt"
	"time"
)

func main() {
	var xd = 1
	var chd = make(chan int)

	go func() {
		chd <- xd         //发送
		fmt.Printf("已发送") //无法执行，缺少接收者而堵塞
	}()

	time.Sleep(1 * time.Second)
}
