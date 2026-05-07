package main

//func main() {
//	ch := make(chan int, 2)
//	go func() {
//		for i := 1; i <= 5; i++ {
//			ch <- i
//			fmt.Println("Отправил:", i)
//			time.Sleep(200 * time.Millisecond)
//		}
//	}()
//	for i := 1; i <= 5; i++ {
//		time.Sleep(1 * time.Second)
//		fmt.Println("Получено:", <-ch)
//	}
//}
