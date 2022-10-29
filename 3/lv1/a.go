package main

//Channel操作代替加锁操作
//缓存为1的channel实现缓存满了阻塞，清空就解除阻塞
import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

var chanRep = make(chan int, 1)

func add() {
	for i := 0; i < 50000; i++ {
		chanRep <- 1 //上锁，因为再来数据就阻塞
		x = x + 1
		<-chanRep //解锁，因为一个缓存被接受了
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
