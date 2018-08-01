
<img src="https://github.com/KenNaNa/go_learing/blob/master/img/42.png"/>

```
package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnter struct {
	name string
}

func (pc PhoneConnter) Name() string {
	return pc.name
}

func (pc PhoneConnter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	var a USB
	a = PhoneConnter{"PhoneConnter"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnect")
}
```
