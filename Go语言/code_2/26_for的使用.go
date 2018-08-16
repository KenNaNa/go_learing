package main //必须有一个main包

import "fmt"

func main() {

	//for 初始化条件 ;  判断条件 ;  条件变化 {
	//}
	//1+2+3 …… 100累加

	sum := 0

	//1) 初始化条件  i := 1
	//2) 判断条件是否为真， i <= 100， 如果为真，执行循环体，如果为假，跳出循环
	//3) 条件变化 i++
	//4) 重复2， 3， 4
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)
}
