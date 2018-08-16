package main //必须有个main包

import "test"
import "fmt"

func main() {
	//包名.函数名
	test.MyFunc()

	//包名.结构体里类型名
	var s test.Stu
	s.Id = 666
	fmt.Println("s = ", s)
}
