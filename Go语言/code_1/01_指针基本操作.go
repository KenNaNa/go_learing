package main //必须有个main包

import "fmt"

func main() {
	var a int = 10
	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a) //变量的内存
	fmt.Printf("&a = %v\n", &a)

	//保存某个变量的地址，需要指针类型   *int 保存int的地址， **int 保存 *int 地址
	//声明(定义)， 定义只是特殊的声明
	//定义一个变量p, 类型为*int
	var p *int
	p = &a //指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 666 //*p操作的不是p的内存，是p所指向的内存(就是a)
	fmt.Printf("*p = %v, a = %v\n", *p, a)

}
