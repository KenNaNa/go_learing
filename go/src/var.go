package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Println(1 ^ 2)
	fmt.Println(!true)
	fmt.Println(1 << 10)
	fmt.Println(1 >> 10)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 &^ 11)
	fmt.Println(6 * 11)
	fmt.Println(6 / 11)

	if a > 0 && (10/a) > 1 {
		fmt.Println("Ok")
	}
}
