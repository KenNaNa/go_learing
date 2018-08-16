package test

import "fmt"

//如果首字母是小写，只能在同一个包里使用
type stu struct {
	id int
}

type Stu struct {
	//id int //如果首字母是小写，只能在同一个包里使用
	Id int
}

//如果首字母是小写，只能在同一个包里使用
func myFunc() {
	fmt.Println("this is myFunc")
}

func MyFunc() {
	fmt.Println("this is MyFunc -=======")
}
