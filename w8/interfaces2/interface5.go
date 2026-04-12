package main

import "fmt"

type Messenger interface {
	Connect() bool
	Send(message string)
	Receive() string
}

type Whatsapp struct {
	isConnected bool
}

func (w *Whatsapp) Connect() bool {
	w.isConnected = true
	return true
}

func (w *Whatsapp) Send(message string) {
	if w.isConnected {
		fmt.Println("Sending from whatsapp:", message)
	} else {
		fmt.Println("No connection")
	}
}

func (w *Whatsapp) Receive() string {
	if w.isConnected {
		return "Aza kalaisyn? Ne janalyk?"
	} else {
		return ""
	}
}

func Chat(msn Messenger) { // msn Messenger = &Whatsapp{}
	msn.Connect()
	msn.Send("Salem dosym")
	fmt.Println(msn.Receive())
}

//func main() {
//	wh := Whatsapp{isConnected: false}
//	Chat(&wh)
//}
