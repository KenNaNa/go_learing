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
