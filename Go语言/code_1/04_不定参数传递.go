package main //必须

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func myfunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func test(args ...int) {
	//全部元素传递给myfunc
	//myfunc(args...)

	//只想把后2个参数传递给另外一个函数使用
	myfunc2(args[:2]...) //args[0]~args[2]（不包括数字args[2]）， 传递过去
	myfunc2(args[2:]...) //从args[2]开始(包括本身)，把后面所有元素传递过去
}

func main() {
	test(1, 2, 3, 4)
}
