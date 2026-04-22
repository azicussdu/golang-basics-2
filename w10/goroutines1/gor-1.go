package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

//func main() {
//	go sayHello()
//
//	fmt.Println("Main function")
//	time.Sleep(1 * time.Second)
//	fmt.Println("After 1 second")
//}
