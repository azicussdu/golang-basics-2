package main

import (
	"fmt"
	"time"
)

func worker(hb chan<- string, shd <-chan bool) {
	for {
		select {
		case <-shd:
			fmt.Println("We are stoping")
			return
		case <-time.After(1 * time.Second):
			hb <- "tuk"
		}
	}
}

//func main() {
//	heartbeat := make(chan string)
//	shutdown := make(chan bool)
//
//	go worker(heartbeat, shutdown)
//
//	timeoutChannel := time.After(5 * time.Second)
//	for {
//		select {
//		case msg := <-heartbeat:
//			fmt.Println(msg)
//		case <-timeoutChannel:
//			shutdown <- true
//			time.Sleep(100 * time.Millisecond)
//			return
//		}
//	}
//}
