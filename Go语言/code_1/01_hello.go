//1) go语言以包作为管理单位
//2) 每个文件必须先声明包
//3) 程序必须有一个main包（重要）
package main

import "fmt"

//入口函数
func main() { //左括号必须和函数名同行
	//打印
	//"hello go"打印到屏幕， Println()会自动换行
	//调用函数，大部分都需要导入包
	/*
		这也是注释， 这是块注释
	*/
	fmt.Println("hello go") //go语言语句结尾是没有分号

	fmt.Println("hello itcast")
}
