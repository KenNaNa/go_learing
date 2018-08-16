package main //必须

import (
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("r = ", calc.Minus(10, 5))
}
