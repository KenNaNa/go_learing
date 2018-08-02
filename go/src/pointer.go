package main

import "fmt"

func main() {
	a := 1
	var p *int = &a
	fmt.Println(p)  //取地址
	fmt.Println(*p) //取内容
	fmt.Println(&a) //取地址
}
