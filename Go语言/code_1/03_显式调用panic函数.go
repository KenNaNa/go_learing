package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
