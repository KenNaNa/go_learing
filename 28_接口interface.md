
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

```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
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
	if pc, ok := usb.(PhoneConnter); ok {
		fmt.Println("Disconnect", pc.name)
		return
	}
	fmt.Println("unknown decive")
}

```

```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
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

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("unknown decive")
	}

}

```

```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
}

func (pc PhoneConnter) Name() string {
	return pc.name
}

func (pc PhoneConnter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	pc := PhoneConnter{"PhoneConnter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("unknown decive")
	}

}
```


接口转换的问题
```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
}
type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("Connected", tv.name)
}
func (pc PhoneConnter) Name() string {
	return pc.name
}

func (pc PhoneConnter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	tv := TVConnecter{"TVConnecter"}
	var a USB
	a = USB(tv)
	a.Connect()
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("unknown decive")
	}

}

```

```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
}

func (pc PhoneConnter) Name() string {
	return pc.name
}

func (pc PhoneConnter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	pc := PhoneConnter{"PhoneConnter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()

	pc.name = "PC"
	a.Connect()
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("unknown decive")
	}

}

```

```
package main

import "fmt"

type empty interface {
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnter struct {
	name string
}

type Connecter interface {
	Connect()
}

func (pc PhoneConnter) Name() string {
	return pc.name
}

func (pc PhoneConnter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	// pc := PhoneConnter{"PhoneConnter"}
	// var a Connecter
	// a = Connecter(pc)
	// a.Connect()

	// pc.name = "PC"
	// a.Connect()
	var a interface{}
	fmt.Println(a == nil) //true

	var p *int = nil
	a = p
	fmt.Println(a == nil)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("unknown decive")
	}

}

```
