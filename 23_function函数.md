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
