package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延时1s
	}
}

func main() {

	go newTask() //新建一个协程， 新建一个任务

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) //延时1s
	}
}
