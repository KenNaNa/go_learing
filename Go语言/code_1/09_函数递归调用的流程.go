package main //必须

import "fmt"

func test(a int) {
	if a == 1 { //函数终止调用的条件，非常重要
		fmt.Println("a = ", a)
		return //终止函数调用
	}

	//函数调用自身
	test(a - 1)

	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
