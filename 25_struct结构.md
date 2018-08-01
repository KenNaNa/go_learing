<img src="https://github.com/KenNaNa/go_learing/blob/master/img/38.png">

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{}
	fmt.Println(a)
}

```

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)
}
```

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{
		Name: "Unamed",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)
}
```


```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{
		Name: "Unamed",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)

	A(a)

	fmt.Println(a)
}

func A(per person) {
	per.Age = 20
	fmt.Println("A", per)
}

```

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{
		Name: "Unamed",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)

	A(&a)

	fmt.Println(a)
}

func A(per *person) {
	per.Age = 20
	fmt.Println("A", per)
}
```

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := person{
		Name: "Unamed",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)

	A(&a)
	B(&a)

	fmt.Println(a)
}

func A(per *person) {
	per.Age = 20
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 18
	fmt.Println("A", per)
}
```

```
package main

import "fmt"

type person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	a := &person{
		Name: "Unamed",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)

	a.Name = "Ken"
	a.Age = 18
	a.Sex = "Man"

	fmt.Println(a)

	A(a)
	B(a)

	fmt.Println(a)
}

func A(per *person) {
	per.Age = 20
	fmt.Println("A", per)
}
```

匿名结构

```
package main

import "fmt"

func main() {
	a := struct {
		Name string
		Age  int
		Sex  string
	}{
		Name: "Ken",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(a)
}

```

```
package main

import "fmt"

func main() {
	a := &struct {
		Name string
		Age  int
		Sex  string
	}{
		Name: "Ken",
		Age:  18,
		Sex:  "Man",
	}
	fmt.Println(*a)
}
```

```
package main

import "fmt"

type person struct {
	Name   string
	Age    int
	Concat struct {
		Phone, City string
	}
}

func main() {
	a := &person{
		Name: "Ken",
		Age:  18,
	}
	a.Concat.City = "广东"
	a.Concat.Phone = "88888888"

	fmt.Println(*a)
}

```

```
package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	a := &person{
		"Ken",
		18,
	}

	fmt.Println(*a)
}
```
下面这种会报错，没有对应上
```
package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	a := &person{
		
		18,
		"Ken",
	}

	fmt.Println(*a)
}
```

```
package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	a := &person{
		"Ken",
		18,
	}
	b := &person{
		"Ken",
		18,
	}

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(a == b)
}
```
```
package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	a := person{
		"Ken",
		18,
	}
	b := person{
		"Ken",
		18,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == b)
}
```
结构嵌入
```
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
		Sex:  0,
	}
	b := student{
		Name: "Ken",
		Age:  18,
		Sex:  1,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == b)
}
```

```
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
		human{
			Sex: 0,
		},
	}

	b := student{
		Name: "Ken",
		Age:  18,
		human{
			Sex: 1,
		},
	}

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(a == b)
}

```


```
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
	// fmt.Println(a == b)
}
```

```
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
```
