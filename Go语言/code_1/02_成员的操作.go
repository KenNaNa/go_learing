package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"

	s1.Person = Person{"go", 'm', 18}

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

}
