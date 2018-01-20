package main

import (
	"fmt"
)

type Connect interface {
	Connect()
}

type Usb interface {
	Name() string
	Connect
}

type PhoneConnect struct {
	name string
}

func (pc *PhoneConnect) Name() string {
	return pc.name
}

func (pc *PhoneConnect) Connect() {
	pc.name = "connc"
	fmt.Println("connect " + pc.name)
}

type TVConnector struct {
}

func main() {

	pc := &PhoneConnect{"phone connect"}
	pc.Connect()

	fmt.Println(pc)
	DisConnectx(pc)

	fmt.Println("===============")

	pc1 := &PhoneConnect{name: "zc"}
	var con Connect
	con = Connect(pc1)
	con.Connect()
	fmt.Println(con)

	var pc2 interface{}

	pc2 = interface{}(pc1)

	fmt.Println("=======switch========")
	switch v := pc2.(type) {
	case PhoneConnect:
		fmt.Println("=====", v.name)
		//fallthrougdh
	case Connect:
		v.Connect()
	default:
		fmt.Println("Unknow")
	}

	fmt.Println("=======copy inteface========")
	pc3 := &PhoneConnect{name: "zc"}
	var con1 Connect
	con1 = Connect(pc3)
	pc3.name = "ppp"
	con1.Connect()
}

func DisConnect(usb Usb) {
	fmt.Println("Disconnected ", usb.Name())
}

func DisConnectx(usb interface{}) {
	//if pc, ok := usb.(*PhoneConnect); ok {
	//	fmt.Println("Match pc", pc.name)
	//	return
	//}

	switch v := usb.(type) {
	case *PhoneConnect:
		fmt.Println("Match pc", v.name)
	default:
		fmt.Println("Unknow")
	}

}
