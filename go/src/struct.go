package main

import "fmt"

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{
		Name: "Ken",
		Age:  18,
		human: human{
			Sex: 0,
		},
	}

	b := student{
		Name: "Ken",
		Age:  18,
		human: human{
			Sex: 1,
		},
	}

	fmt.Println(a)
	fmt.Println(b)

	a.Name = "NaNa"
	a.Age = 15
	a.Sex = 100

	fmt.Println(a)
}
