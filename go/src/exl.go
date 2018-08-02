package main

import "fmt"

const (
	B  float64 = 1 << (iota * 10)
	KB float64 = 1 << (iota * 10)
	MB float64 = 1 << (iota * 10)
	GB float64 = 1 << (iota * 10)
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
