package main

import "fmt"

// 切片

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 := arr[5:7]                       //[5 6]
	s3 := append(s2, 10)                 //[5 6 10]
	s4 := append(s3, 11)                 //[5 6 10 11]
	s5 := append(s4, 12)                 // [5 6 10 11 12]
	fmt.Println("s3,s4,s5=", s3, s4, s5) //s3,s4,s5= [5 6 10] [5 6 10 11] [5 6 10 11 12]
	fmt.Println("arr=", arr)             //arr= [0 1 2 3 4 5 6 10]
}
