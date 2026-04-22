package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	start := time.Now()
	sum := 0

	for i := 0; i < 100_000_000; i++ {
		sum += i
	}

	elapsed := time.Since(start)
	fmt.Printf("Горутина %d завершилась за %v\n", id, elapsed)
}

//func main() {
//	runtime.GOMAXPROCS(4) // задать количество потоков ОС
//
//	for i := 1; i <= 4; i++ {
//		go worker(i)
//	}
//
//	time.Sleep(5 * time.Second)
//}
