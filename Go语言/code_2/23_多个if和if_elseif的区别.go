package main //必须有一个main包

import "fmt"

func main() {
	//这种好
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}

	if b > 10 {
		fmt.Println("b > 10")
	}

	if b < 10 {
		fmt.Println("b < 10")
	}

}
