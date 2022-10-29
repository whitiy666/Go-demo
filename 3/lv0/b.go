package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func bar1(dishes string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("正在烤"+dishes+",熟度为:%d分熟\n", i*2)
		time.Sleep(time.Second)
	}

	fmt.Printf(dishes + "烤熟了！")
	wg.Done() //用于减少协程数据
}

func main() {
	wg.Add(2) //添加两个协程数据，要用到
	go bar1("学长")
	go bar1("学弟")
	wg.Wait() //避免一个协程运行完了就退出

	fmt.Printf("退出主协程")

}
