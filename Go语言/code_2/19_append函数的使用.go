package main //必须有个main包

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2 = ", s2)
}
