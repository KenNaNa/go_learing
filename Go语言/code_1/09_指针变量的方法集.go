package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	//p.SetInfoPointer() //func (p *Person) SetInfoPointer()
	(*p).SetInfoPointer() //把(*p)转换层p后再调用，等价于上面

	//内部做的转换， 先把指针p， 转成*p后再调用
	//(*p).SetInfoValue()
	//p.SetInfoValue()

}
