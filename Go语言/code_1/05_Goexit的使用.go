package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccccc")

	//return //终止此函数
	runtime.Goexit() //终止所在的协程

	fmt.Println("dddddddddddddddddddddd")
}

func main() {

	//创建新建的协程
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaa")

		//调用了别的函数
		test()

		fmt.Println("bbbbbbbbbbbbbbbbbbb")
	}() //别忘了()

	//特地写一个死循环，目的不让主协程结束
	for {
	}
}
