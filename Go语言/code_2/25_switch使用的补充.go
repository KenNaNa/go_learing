package main //必须有一个main包

import "fmt"

func main() {
	//支持一个初始化语句， 初始化语句和变量本身， 以分号分隔
	switch num := 4; num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")

	case 2:
		fmt.Println("按下的是2楼")

	case 3, 4, 5:
		fmt.Println("按下的是yyy楼")

	case 6:
		fmt.Println("按下的是4楼")

	default:
		fmt.Println("按下的是xxx楼")
	}

	score := 85
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80: //case后面可以放条件
		fmt.Println("良好")
	case score > 70: //case后面可以放条件
		fmt.Println("一般")
	default:
		fmt.Println("其它")
	}
}
