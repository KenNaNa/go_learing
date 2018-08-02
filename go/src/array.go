package main

import "fmt"

func main() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)

	num := len(a)

	for i := 0; i < num; i++ {
		v := 1
		fmt.Println(&v)
	}
}
