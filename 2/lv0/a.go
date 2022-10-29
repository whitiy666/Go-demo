package main

import "fmt"

type Student struct {
	Name  string
	Class int
	Age   int8
}

type Person struct {
	Name   string
	Sex    string
	Nation string
}

type Chinese struct {
	Basic   Person
	Address string
}

func main() {
	xiaoMing := Student{
		Name:  "小明",
		Class: 114514,
		Age:   3,
	}

	var xiaoHua Student
	xiaoHua.Name = "小花"
	xiaoHua.Class = 520114
	xiaoHua.Age = 11

	var leLe = Chinese{
		Basic: Person{
			Sex:    "man",
			Nation: "cn",
			Name:   "乐乐",
		},
		Address: "chongqing",
	}

	fmt.Printf("姓名：%s，年龄：%d，班级：%d\n", xiaoMing.Name, xiaoMing.Age, xiaoMing.Class)
	fmt.Printf("姓名：%s，年龄：%d，班级：%d\n", xiaoHua.Name, xiaoHua.Age, xiaoHua.Class)
	fmt.Printf("姓名：%s，性别:%s，地址:%s，国籍:%s\n", leLe.Basic.Name, leLe.Basic.Sex, leLe.Address, leLe.Basic.Nation)

	fmt.Printf("人们要做作业了\n")
	leLe.Homework()

}

func (p Chinese) Homework() error {
	fmt.Printf("%s正在做作业", p.Basic.Name)
	return nil
}
