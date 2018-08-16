package main //必须有一个main包

import "fmt"

//go函数可以返回多个值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//	a := 10
	//	b := 20
	//	c := 30

	a, b := 10, 20

	//交换2个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b = %d\n", a, b)

	//	i := 10
	//	j := 20
	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	i = 10
	j = 20

	//_匿名变量，丢弃数据不处理, _匿名变量配合函数返回值使用，才有优势
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var c, d, e int
	c, d, e = test() //return 1, 2, 3
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	_, d, e = test() //return 1, 2, 3
	fmt.Printf("d = %d, e = %d\n", d, e)

}
