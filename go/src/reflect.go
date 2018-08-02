package main

import "fmt"

import "reflect"

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("hello", name, "my name is", u.Name)
}
func main() {
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("Ken")}
	mv.Call(args)
}
