package main //必须

import "fmt"

type FuncType func(int, int) int

//实现加法
func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
//现有想法，后面再实现功能
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b) //这个函数还没有实现
	//result = Add(a, b) //Add()必须先定义后，才能调用
	return
}

func main() {
	a := Calc(1, 1, Mul)
	fmt.Println("a = ", a)
}
