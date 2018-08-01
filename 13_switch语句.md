<img src="https://github.com/KenNaNa/go_learing/raw/master/img/21.png"/>

```
package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("None")
	}
}
```

```
package main

import "fmt"

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a = 0")
		fallthrough
	case a >= 1:
		fmt.Println("a = 1")
		fallthrough
	case a >= 2:
		fmt.Println("a = 2")
		fallthrough
	default:
		fmt.Println("None")
	}
}
```

```
package main

import "fmt"

func main() {

	switch a := 1; {
	case a >= 0:
		fmt.Println("a = 0")
		fallthrough
	case a >= 1:
		fmt.Println("a = 1")
	case a >= 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("None")
	}
}
```
