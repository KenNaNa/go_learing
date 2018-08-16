package main //必须有一个main包

import "fmt"

func main() {
	//1
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//2
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//3
	a = 8
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

	//4
	if a := 8; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

}
