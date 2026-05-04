package main

import (
	"fmt"
	"time"
)

func sayHello(letter string) { // letter = A, B, C
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from ", letter)
		time.Sleep(100 * time.Millisecond)
	}
}

//func main() {
//	go sayHello("A")
//	go sayHello("B")
//	go sayHello("C")
//
//	time.Sleep(1 * time.Second)
//}
