package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	Person //结构体匿名字段
	int    //基础类型的匿名字段
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "hehehe"}
	fmt.Printf("s = %+v\n", s)

	s.Person = Person{"go", 'm', 22}

	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person, s.int, s.mystr)

}
