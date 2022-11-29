package main

import "fmt"

// 删除slice头尾
func main() {
	s1 := []int{1, 2, 3, 4, 5, 0, 0, 0}
	fmt.Println(s1) //[1 2 3 4 5 0 0 0]

	//删除头
	front := s1[0] //1
	s1 = s1[1:]

	//删除尾
	tail := s1[len(s1)-1] //0
	s1 = s1[:len(s1)-1]

	fmt.Println(front, tail) //1 0
	fmt.Println(s1)          //[2 3 4 5 0 0]
}
