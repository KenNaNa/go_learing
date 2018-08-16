package main //必须

import "fmt"

func main() {
	a := 10
	str := "mike"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}() //()代表直接调用

	fmt.Printf("外部：a = %d, str = %s\n", a, str)

}
