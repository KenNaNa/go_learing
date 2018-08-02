package main

import "fmt"

// import "runtime"
// import "sync"
var c chan string

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("format", i)
		i++
	}
}
func main() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("format", i)
		fmt.Println(<-c)
	}
}
