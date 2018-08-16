package main //必须有一个main包

import "fmt"

func main() {

	//break //break is not in a loop, switch, or select
	//continue//continue is not in a loop

	//goto可以用在任何地方，但是不能夸函数使用
	fmt.Println("11111111111111")

	goto End //goto是关键字， End是用户起的名字， 他叫标签

	fmt.Println("222222222222222")

End:
	fmt.Println("3333333333333")

}
