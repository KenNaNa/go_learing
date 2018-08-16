package main //必须

import "fmt"

var a byte //全局变量

func main() {
	var a int //局部变量

	//1、不同作用域，允许定义同名变量
	//2、使用变量的原则，就近原则
	fmt.Printf("%T\n", a) //int

	{
		var a float32
		fmt.Printf("2: %T\n", a) //float32
	}

	test()

}

func test() {
	fmt.Printf("3: %T\n", a) //uint8, 就是byte
}
