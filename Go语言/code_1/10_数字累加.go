package main //必须

import "fmt"

//实现1+2+3+……100
func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}

	return
}

func test02(i int) int {
	if i == 1 {
		return 1
	}

	return i + test02(i-1)
}

func test03(i int) int {
	if i == 100 {
		return 100
	}

	return i + test03(i+1)
}

func main() {
	var sum int
	//sum = test01()
	//sum = test02(100)
	sum = test03(1)
	fmt.Println("sum = ", sum)

}
