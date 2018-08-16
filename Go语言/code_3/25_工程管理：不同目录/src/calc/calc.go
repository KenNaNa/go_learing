package calc

import "fmt"

func init() {
	fmt.Println("this is calc init")
}

//func add(a, b int) int {
func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
