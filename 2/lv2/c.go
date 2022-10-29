package main

import "fmt"

type Sayer interface {
	Say() //接口
}

func AnmSay(s Sayer) {
	switch s.(type) {
	case Dog:
		fmt.Printf("汪汪汪\n") //实现狗
		break
	case Cat:
		fmt.Printf("喵喵喵\n") //实现猫
		break
	default:
		fmt.Printf("不知道怎么叫为好了\n")
		break

	}
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Say() {
	fmt.Printf("%s汪汪汪\n", d.Name) //实现狗
}
func (c Cat) Say() {
	fmt.Printf("%s喵喵喵\n", c.Name) //实现猫
}

func main() {
	Mao := Cat{Name: "猫猫"}
	Gou := Dog{Name: "狗狗"}

	Mao.Say()
	Gou.Say()

	AnmSay(Mao)
	AnmSay(Gou)

}
