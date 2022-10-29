package main

import (
	"fmt"
)

func main() {
	fmt.Println("输入原始数据（string）:")
	var originData string
	fmt.Scanln(&originData) //读取用户输入，储存到od变量

	data := make([]rune, 0, len(originData)) //生成切片

	for i := 0; i < len(data); i++ {
		if data[i] != data[len(data)-i-1] {
			fmt.Println("false")
		}

	} //for循环判断，正数和倒数是一致的回文
	output := originData[0 : len(originData)/2]
	fmt.Println("回文：", output)

}
