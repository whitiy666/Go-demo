package main

import "fmt"

// 验证 在能够接受后不会堵塞
func main() {
	var channel = make(chan int)
	var send = 6666
	var receive int
	go func() {
		channel <- send
		fmt.Println("已发送")
	}()
	// 使用receive接受从channel中传来的值
	receive = <-channel //接受了就可以不堵塞了
	fmt.Println(receive)
}
