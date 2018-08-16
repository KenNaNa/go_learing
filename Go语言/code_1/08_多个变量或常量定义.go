package main //必须有一个main包

import "fmt"

func main() {
	//不同类型变量的声明(定义)
	//	var a int = 1
	//	var b float64 = 2.0
	//	var (
	//		a int     = 1
	//		b float64 = 2.0
	//	)

	//可以自动推导类型
	var (
		a = 1
		b = 2.0
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	//选中代码，按ctrl+\， 注释和取消注释代码的快捷键
	//	const i int = 10
	//	const j float64 = 3.14

	//	const (
	//		i int     = 10
	//		j float64 = 3.14
	//	)

	//可以自动推导类型
	const (
		i = 10
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
