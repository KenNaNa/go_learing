package main //必须有一个main包

import "fmt"
import "time"

func main() {

	i := 0

	for { //for后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second) //演示1s

		if i == 5 {
			//break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			continue //跳过本次循环，下一次继续
		}
		fmt.Println("i = ", i)
	}
}
