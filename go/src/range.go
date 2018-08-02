package main

import "fmt"
import "sort"

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)
	m2:=make(map[string][int])
	for k,v:=range m1{
		m2[v] = k
	}
	fmt.Println(m2)

}
