package main

import "fmt"

type Dog struct {
	Name string
	Age  int
} //类型狗

func (d Dog) Wang() {
	fmt.Printf("%s:汪汪汪", d.Name)
} //狗的狗叫方法

func main() {
	keke := Dog{
		Name: "可可",
		Age:  1,
	} //定义可可狗
	keke.Wang() //可可狗叫
}
