package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("plan.txt") //创建

	if err != nil {
		fmt.Printf("创建失败")

	} else {
		data := "I’m not afraid of difficulties and insist on learning programming"
		_, err = f.Write([]byte(data)) //写入
		if err != nil {
			fmt.Printf("写入失败")
		}
	}

	//读取部分

	//file, err := os.Open("plan.txt")
	//if err != nil {
	//	fmt.Printf("读取失败")
	//}
	//
	//先鸽，读取到缓冲区，大概要把缓冲区内容for无限（还要在末尾-读取出错就是末尾大概，break掉循环）读取出来，最后打印缓冲区，有空再看吧。按道理这种操作应该有现成的库偷懒

	//19：40 2022，10，29吃饭去了

}
