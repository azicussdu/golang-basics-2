package main

import (
	"fmt"
	"time"
)

func send5(ch chan string) { // kanalga 5 soz jazady, sonynda close()
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("message %d\n", i)
		time.Sleep(400 * time.Millisecond)
	}
	close(ch)
}

//func main() {
//	ch := make(chan string)
//	go send5(ch)
//
//	for {
//		msg, ok := <-ch
//		if !ok { // if channel is closed
//			fmt.Println("Channel is closed")
//			break
//		}
//		fmt.Println("Received:", msg)
//	}
//
//	fmt.Println("Program finished")
//}
