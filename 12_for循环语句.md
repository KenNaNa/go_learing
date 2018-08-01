<img src="https://github.com/KenNaNa/go_learing/blob/master/img/20.png"/>

```
package main

import "fmt"

func main() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
		// 表示无限循环
	}
	fmt.Println("over")
}
```


```
package main

import "fmt"

func main() {
	a := 1
	for a <= 3 {
		a++
		fmt.Println(a)
		// 表示无限循环
	}
	fmt.Println("over")
}
```

```
package main

import "fmt"

func main() {
	a := 1
	for i := 1; i < 3; i++ {
		a++
		fmt.Println(a)
		// 表示无限循环
	}
	fmt.Println("over")
}
```

```
package main

import "fmt"

func main() {
	a := 1
	b := "string"
	l := len(b)
	for i := 1; i < l; i++ {
		a++
		fmt.Println(a)
		// 表示无限循环
	}
	fmt.Println("over")
}
```
