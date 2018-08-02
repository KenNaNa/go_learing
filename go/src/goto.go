package main

import "fmt"

func main() {
Label1:
	for i := 0; i < 10; i++ {
		for {
			// 跳出整级循环
			fmt.Println(i)
			continue Label1
		}
	}
	fmt.Println("OK")
}
