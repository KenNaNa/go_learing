package main //必须

import "fmt"

//多个返回值
func myfunc01() (int, int, int) {
	return 1, 2, 3
}

//go官方推荐写法
func myfunc02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}

func myfunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	//函数调用
	a, b, c := myfunc02()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
