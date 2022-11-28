package main

import "fmt"

// 指针
// 不能运算
func main() {
	var a int = 2
	var pa *int = &a
	fmt.Println(pa)
	*pa = 3
	fmt.Println(a)
}
