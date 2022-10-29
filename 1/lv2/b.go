package main

import "fmt"

var valueOld int
var index = 0 //0:初次使用 ， 1：进行了计算

func main() {
	var op string
	var v1, v2 int

	fmt.Println("输入计算式子，如 1 + 1,如果要退出就直接结束进程吧，：")
	fmt.Scanln(&v1, &op, &v2)
	operate(v1, v2, op)

}

func operate(v1, v2 int, op string) {
	switch op {
	case "+":
		fmt.Println(add(v1, v2))
		fmt.Scanln(&v1, &op, &v2) //获取新的表达式
		operate(v1, v2, op)

	case "-":
		fmt.Println(reduce(v1, v2))
		fmt.Scanln(&v1, &op, &v2) //获取新的表达式
		operate(v1, v2, op)
	case "÷":
		fmt.Println(divide(v1, v2))
		fmt.Scanln(&v1, &op, &v2) //获取新的表达式
		operate(v1, v2, op)
	case "/":
		fmt.Println(divide(v1, v2))
		fmt.Scanln(&v1, &op, &v2) //获取新的表达式
		operate(v1, v2, op)
	case "*":
		fmt.Println(multiply(v1, v2))
		fmt.Scanln(&v1, &op, &v2) //获取新的表达式
		operate(v1, v2, op)

	default:
		fmt.Println("无法运算，结束")

	}
}

func add(v1, v2 int) int {

	if index == 0 {
		valueOld = v1 + v2
		index = 1
		return v1 + v2
	} else {
		fmt.Println(valueOld)
		valueOld = v1 + v2
		return v1 + v2
	}

}

func reduce(v1, v2 int) int { //减法
	if index == 0 {
		valueOld = v1 - v2
		index = 1
		return v1 - v2
	} else {
		fmt.Println(valueOld)
		valueOld = v1 - v2
		return v1 - v2
	}
}

func multiply(v1, v2 int) int { //乘法
	if index == 0 {
		valueOld = v1 * v2
		index = 1
		return v1 * v2
	} else {
		fmt.Println(valueOld)
		valueOld = v1 * v2
		return v1 * v2
	}

}

func divide(v1, v2 int) int { //除法
	if index == 0 {
		valueOld = v1 / v2
		index = 1
		return v1 / v2
	} else {
		fmt.Println(valueOld)
		valueOld = v1 / v2
		return v1 / v2
	}
}
