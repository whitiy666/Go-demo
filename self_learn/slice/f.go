package main

import "fmt"

// 删除切片的值
func main() {
	s1 := []int{1, 2, 3, 4, 5, 0, 0, 0}
	fmt.Println(s1)

	s1 = append(s1[:3], s1[4:]...)
	fmt.Println(s1)
}
