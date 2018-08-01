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
