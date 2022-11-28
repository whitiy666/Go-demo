package main

import "fmt"

// 交换值 法一
func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	a, b := 3, 4
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
