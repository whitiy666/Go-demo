package main

import (
	"fmt"
	"sync"
)

//锁🔒 保证只有一个gor.t.访问数据，避免数据竞争的问题

var x int64

var wgc sync.WaitGroup
var mu sync.Mutex

func main() {

	wgc.Add(2)
	go add()
	go add()
	wgc.Wait() //等待子协程全部完毕
	fmt.Printf("结果：%d，退出主协程", x)

}

func add() {
	for i := 0; i < 5000; i++ {
		mu.Lock() //数据上锁
		x++
		mu.Unlock() //解锁
	}
	wgc.Done() //减少协程
}
