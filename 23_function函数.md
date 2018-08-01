<img src="https://github.com/KenNaNa/go_learing/blob/master/34.png">
<img src="https://github.com/KenNaNa/go_learing/blob/master/33.png">
<img src="https://github.com/KenNaNa/go_learing/blob/master/32.png">

```
package main

import "fmt"

func main() {

}

func A() (a b c int) {
	a, b, c = 1, 2, 4
	return a, b, c
}

func B(a, b, c int) (int, int, int) {
	return a,b,c
}

func C(a ...int) {

}

```

```
package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 5}
	A(s)
	fmt.Println(s)
}

func A(s []int) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}
```

```
package main

import "fmt"

func main() {
	a := 1
	A(&a)
	fmt.Println(a)
}

func A(a *int) {
	*a = 2
	fmt.Println(*a)
}
```

```
package main

import "fmt"

func main() {
	a := 1
	c := A
	c(&a)
	A(&a)
	fmt.Println(a)
}

func A(a *int) {
	*a = 2
	fmt.Println(*a)
}
```
```
package main

import "fmt"

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

```
<img src="https://github.com/KenNaNa/go_learing/blob/master/35.png">

```
package main

import "fmt"

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
```

```
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
```

这他妈跟javascript的闭包太像了
defer就像那个按钮一样

```
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
```

```
package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("func A")
}
func B() {
	panic("panic in B")
}

func C() {
	fmt.Println("func C")
}
```

```
package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("func A")
}
func B() {
	panic("panic in B")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
}

func C() {
	fmt.Println("func C")
}
```


```
package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")

}

func C() {
	fmt.Println("func C")
}

```
