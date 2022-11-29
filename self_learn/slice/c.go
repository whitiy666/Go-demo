package main

import "fmt"

func printSlice_c(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
func main() {
	s1 := []int{2, 4, 6, 8}
	printSlice_c(s1)
}
