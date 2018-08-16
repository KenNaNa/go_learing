package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1 = ", s1)

	//自动推导类型
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	//fmt.Println("s2 = ", s2)
	//%+v, 显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	//指定成员初始化，没有初始化的常用自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Printf("s4 = %+v\n", s4)

	//s5 := Student{"mike", 'm', 18, 1, "bj"} //err

}
