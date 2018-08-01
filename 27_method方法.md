<img src="https://github.com/KenNaNa/go_learing/blob/master/40.png">

```
package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
}

func (a A) Print() {
	fmt.Println("A")
}
```

就我的理解，可以理解成命名空间吧
```
package main

import "fmt"

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	b := B{}
	b.Print()
}

func (a A) Print() {
	fmt.Println("A")
}

func (b B) Print() {
	fmt.Println("B")
}
```

指针调用
```
package main

import "fmt"

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	b := B{}
	b.Print()
	fmt.Println(b.Name)
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
```

```
package main

import "fmt"

type A int

func main() {
	var a A
	a.Print()
}

func (a *A) Print() {
	fmt.Println("A")
}
```

```
package main

import "fmt"

type A int

func main() {
	var a A
	a.Print()
	(*A).Print(&a)

	// fmt.Println(*A == a)
}

func (a *A) Print() {
	fmt.Println("A")
}
```


```
package main

import "fmt"

type A struct {
	// Name 为导出字段
	// name 为不可导出字段
	name string
}

func main() {
	var a A
	a.Print()
	fmt.Println(a.name)

	// fmt.Println(*A == a)
}

func (a *A) Print() {
	a.name = "124"
	fmt.Println(a.name)
}

```
