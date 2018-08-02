package main

import "fmt"

func main() {
	a := 1
	b := "string"
	l := len(b)
	for i := 1; i < l; i++ {
		a++
		fmt.Println(a)
		// 表示无限循环
	}
	fmt.Println("over")
}
