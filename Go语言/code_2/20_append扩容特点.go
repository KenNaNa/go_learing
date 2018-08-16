package main //必须有个main包

import "fmt"

func main() {
	//如果超过原来的容量，通常以2倍容量扩容
	s := make([]int, 0, 1) //容量为1
	oldCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d ===> %d\n", oldCap, newCap)

			oldCap = newCap
		}
	}

}
