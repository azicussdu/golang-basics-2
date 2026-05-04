package main

//func main() {
//	ch := make(chan int)
//	quit := make(chan bool)
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch <- i // отправляем значения в канал ch
//			time.Sleep(time.Second)
//		}
//		close(ch) // ну и в конце закроем канал
//	}()
//
//	go func() {
//		for {
//			select {
//			case value, ok := <-ch:
//				if !ok {
//					fmt.Println("Канал закрыт, выхожу:", value)
//					quit <- true
//					return
//				}
//				fmt.Printf("Получил: %d\n", value)
//			case <-time.After(500 * time.Millisecond):
//				fmt.Println("Простой на 500ms")
//			}
//		}
//	}()
//
//	<-quit
//	fmt.Println("Завершение программы")
//}
