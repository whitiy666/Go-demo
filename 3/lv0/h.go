package main

import "fmt"

// 通道死锁
func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}

func f1(in chan int) {
	fmt.Println(<-in)
}

//主协程和协程都被阻塞了（因为主线程中的chan期待一个接收它的，而协程中期待一个发送它的）。所以产生通道死锁的原因是通道out写入2,这个时候是同步等待接收端接收的状态,而接收的代码却在下一行,这个时候就导致了整个程序无法往下执行,造成死锁。
