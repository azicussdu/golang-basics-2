package main

import "fmt"

type Transport interface {
	Speed() int
}

func PrintSpeed(trn Transport) { // trn Transport = Car{}
	fmt.Println("This transport speed is:", trn.Speed())
}

type Car struct {
	isTurbo bool
}

func (c Car) Speed() int {
	if c.isTurbo {
		return 320
	} else {
		return 260
	}
}

//func main() {
//	cr := Car{isTurbo: false}
//	PrintSpeed(cr)
//
//}
