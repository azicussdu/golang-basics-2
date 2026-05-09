package main

import (
	"fmt"
	"time"
)

func goworker(id int, semaphore chan struct{}) {
	semaphore <- struct{}{}
	fmt.Println("Jumys jazap jatyr: ", id)
	time.Sleep(1 * time.Second)
	fmt.Println("Jumys bitti: ", id)
	<-semaphore
}

//func main() {
//	sem := make(chan struct{}, 2)
//
//	for i := 1; i <= 8; i++ {
//		go goworker(i, sem)
//	}
//
//	time.Sleep(5 * time.Second)
//}
