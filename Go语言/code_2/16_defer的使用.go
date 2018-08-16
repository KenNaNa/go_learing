package main //必须

import "fmt"

func main() {
	//defer延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbbbbbbbbb")

	fmt.Println("aaaaaaaaaaaaaaa")
}
