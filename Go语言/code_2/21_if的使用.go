package main //必须有一个main包

import "fmt"

func main() {
	s := "屌丝"

	//if和{就是条件，条件通常都是关系运算符
	if s == "王思聪" { //左括号和if在同一行
		fmt.Println("左手一个妹子，右手一个大妈")
	}

	//if支持1个初始化语句, 初始化语句和判断条件以分号分隔
	if a := 10; a == 10 { //条件为真，指向{}语句
		fmt.Println("a == 10")
	}
}
