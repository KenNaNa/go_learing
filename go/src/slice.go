package main

import "fmt"

func main() {
	s1 := []int{1, 2, 4, 5, 6}
	s2 := s1[:5]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(&s1)
	fmt.Println(&s2)

}
