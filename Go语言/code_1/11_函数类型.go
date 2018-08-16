package main //必须

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型， 通过type给一个函数类型起名
//FuncType它是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}

func main() {
	var result int
	result = Add(1, 1) //传统调用方式
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = Add            //是变量就可以赋值
	result = fTest(10, 20) //等价于Add(10, 20)
	fmt.Println("result2 = ", result)

	fTest = Minus
	result = fTest(10, 5) //等价于Minus(10, 5)
	fmt.Println("result3 = ", result)
}
