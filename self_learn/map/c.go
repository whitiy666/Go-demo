package main

import "fmt"

// 获取map的值
func main() {
	m1 := map[string]string{
		"name": "wh",
		"city": "sgp",
	} //map[city:sgp name:wh]

	_, nameok := m1["name"]
	fmt.Println(nameok)

	_, namaok := m1["nama"]
	fmt.Println(namaok)

	if name, ok := m1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("空值")
	}

}
