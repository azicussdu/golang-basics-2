package main

//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//
//	go func() {
//		for {
//			ch1 <- "ping"
//			time.Sleep(1 * time.Second)
//		}
//	}()
//
//	go func() {
//		for {
//			ch2 <- "pong"
//			time.Sleep(1500 * time.Millisecond)
//		}
//	}()
//
//	timeoutChannel := time.After(5 * time.Second)
//
//	for {
//		select {
//		case msg1 := <-ch1:
//			fmt.Println(msg1)
//		case msg2 := <-ch2:
//			fmt.Println(msg2)
//		case <-timeoutChannel:
//			fmt.Println("Program finished")
//			return
//		}
//	}
//}
