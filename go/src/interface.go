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
