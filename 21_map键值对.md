<img src="https://github.com/KenNaNa/go_learing/blob/master/img/31.png">

```
package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)

	s := make(map[int]string)
	fmt.Println(s)
	fmt.Println(m)
}


```


```
package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1] = "Ken"
	a := m[1]
	fmt.Println(a)
	s := make(map[int]string)
	fmt.Println(s)
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(a)
	fmt.Println(m)
}
```

```
package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	m[1] = make(map[int]string)
	m[1][1] = "OK"
	a := m[1][1]
	fmt.Println(a)
	fmt.Println(m)
}
```

```
package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]
	fmt.Println(a, ok)
	if !ok {
		m[2] = make(map[int]string)
		m[2][1] = "Ok"
	}
	a, ok = m[2][1]
	fmt.Println(a, ok)
}
```

```
package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)
}
```

```
package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5)
	for i, _ := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
```

```
package main

import "fmt"
import "sort"

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
```
