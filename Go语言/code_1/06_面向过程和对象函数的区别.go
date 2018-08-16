package main

import "fmt"

//实现2数相加
//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

//tmp叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add01(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)

	//定义一个变量
	var a long = 2
	//调用方法格式： 变量名.函数(所需参数)
	r := a.Add02(3)
	fmt.Println("r = ", r)

	//面向对象只是换了一种表现形式

}
