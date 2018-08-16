package main //必须有个main包

import "fmt"

//数组做函数参数，它是值传递
//实参数组的每个元素给形参数组拷贝一份
//形参的数组是实参数组的复制品
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a = ", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5} //初始化

	modify(a) //数组传递过去
	fmt.Println("main: a = ", a)
}
