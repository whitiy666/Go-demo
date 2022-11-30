package main

import (
	"fmt"
)

//遍历
func main() {
	m1 := map[string]string{
		"name": "wh",
		"city": "sgp",
	} //map[city:sgp name:wh]

	for k, v := range m1 {
		fmt.Println(k, v)

	}

}