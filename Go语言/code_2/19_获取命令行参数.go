package main //必须

import "fmt"
import "os"

func main() {
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	//xxx.exe a b
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
