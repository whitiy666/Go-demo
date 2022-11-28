package main

import "fmt"

//交换值-法二-返回

func swap1(a, b int) (int, int) {
	return b, a
}
func main() {
	a, b := 3, 4
	fmt.Println(a, b)
	a, b = swap1(a, b)
	fmt.Println(a, b)
}
