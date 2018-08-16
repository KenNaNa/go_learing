package main //必须有个main包

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func test02(p *Student) {
	p.id = 666
}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test02(&s) //地址传递（引用传递），形参可以改实参
	fmt.Println("main: ", s)

}

func test01(s Student) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main01() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test01(s) //值传递，形参无法改实参
	fmt.Println("main: ", s)

}
