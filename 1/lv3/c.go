package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var data = make([]int, 100)

	for i := 0; i < 100; i++ {
		data[i] = rand.Intn(100) //姑且给个范围吧

	}
	fmt.Print("原始数据", data)

	//sort.Ints(data)
	order(data)
	fmt.Println("\n")
	fmt.Print("排序数据", data)
}

func order(data []int) []int {

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j+1], data[j] = data[j], data[j+1]
			}

		}
	}
	return data

}
