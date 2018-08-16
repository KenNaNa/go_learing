package main //必须

import "fmt"

//函数定义
func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return //有返回值的函数，必须通过return返回
}

func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)

	//通过匿名变量丢弃某个返回值
	a, _ := MaxAndMin(10, 20)
	fmt.Printf("a = %d\n", a)
}
