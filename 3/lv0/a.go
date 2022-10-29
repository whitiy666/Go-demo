package main

import (
	"fmt"
	"time"
)

func main() {
	go bar("学长")
	bar("学弟")
}

func bar(dishes string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("正在烤"+dishes+",熟度为:%d分熟\n", i*2)
		time.Sleep(time.Second) // 睡眠一秒容易看到效果
	}
	fmt.Println(dishes + "哈哈哈烤好辣!")
} //出现有 学长或学弟没烤好就结束的问题。接下来使用waitgroup解决。
