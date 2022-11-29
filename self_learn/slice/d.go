package main

import "fmt"

func printSlice_d(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
func main() {
	s1 := make([]int, 16)
	s2 := make([]int, 10, 32)
	printSlice_d(s1) //len=16,cap=16
	printSlice_d(s2) //len=10,cap=32
}
