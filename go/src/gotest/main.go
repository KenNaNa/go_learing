package gotest

import "fmt"

func Hello() {
	fmt.Println("hello")
	Hi()
}
func Hi() {
	fmt.Println("hi from gotest")
}
