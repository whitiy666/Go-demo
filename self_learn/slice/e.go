package main

import "fmt"

// copy slice
func main() {
	s1 := []int{1, 2, 3, 4} //[1 2 3 4]
	s2 := make([]int, 8)    //[0 0 0 0 0 0 0 0]
	copy(s2, s1)
	fmt.Println(s2) //[1 2 3 4 0 0 0 0]

}
