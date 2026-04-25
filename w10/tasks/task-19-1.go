package main

import (
	"fmt"
	"sync"
	"time"
)

func printNums(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i++ {
		fmt.Println(num + i)
		time.Sleep(1 * time.Second)
	}
}

//func main() {
//	var wg sync.WaitGroup
//
//	for i := 10; i <= 50; i += 10 {
//		wg.Add(1)
//		go printNums(i, &wg) // 10, 20, 30, 40, 50
//	}
//
//	wg.Wait()
//	fmt.Println("Finishing counting")
//}
