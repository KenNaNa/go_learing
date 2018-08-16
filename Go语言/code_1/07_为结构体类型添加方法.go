package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//带有接收者的函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

type long int

//只要接收者类型不一样，这个方法就算同名，也是不同方法，不会出现重复定义函数的错误
func (tmp long) test() {

}

type char byte

func (tmp char) test() {

}

func (tmp *long) test02() {

}

type pointer *int

//pointer为接收者类型，它本身不能是指针类型
//func (tmp pointer) test() {

//}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()

}
