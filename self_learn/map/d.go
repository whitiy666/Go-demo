package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name": "wh",
		"city": "sgp",
	} //map[city:sgp name:wh]

	fmt.Println(m1)
	delete(m1, "name")
	fmt.Println(m1)
}
