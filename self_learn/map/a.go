package main

import "fmt"

// map定义
func main() {
	m1 := map[string]string{
		"name": "wh",
		"city": "sgp",
	} //map[city:sgp name:wh]

	m2 := make(map[string]int) //map[]

	var m3 map[string]int //map[]

	fmt.Println(m1, m2, m3)
}
