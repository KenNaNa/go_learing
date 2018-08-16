package main

import (
	"fmt"
	"time"
)

//验证time.NewTimer()，时间到了，只会响应一次
func main() {
	timer := time.NewTimer(1 * time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

func main01() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C //channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
