package main //必须有一个main包

import "fmt"

func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a: ")

	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a) //别忘了&
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
