package main //必须有个main包

import "fmt"

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	test(m) //在函数内部删除某个key

	fmt.Println("m = ", m)
}
