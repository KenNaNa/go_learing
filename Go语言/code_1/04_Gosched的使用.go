package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}

	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}
