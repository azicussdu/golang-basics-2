package main

import (
	"fmt"
	"sync"
)

func sumOfPart(data []int, start, end int, totalSum *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	sum := 0
	for i := start; i < end; i++ {
		sum += data[i]
	}

	mu.Lock()
	*totalSum += sum
	mu.Unlock()

	fmt.Println("gr sum =", sum)
}

//func main() {
//	data := make([]int, 1000)
//	for i := 0; i < 1000; i++ {
//		data[i] = rand.IntN(1000)
//	}
//
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	totalSum := 0
//
//	wg.Add(4)
//	go sumOfPart(data, 0, 250, &totalSum, &wg, &mu)
//	go sumOfPart(data, 250, 500, &totalSum, &wg, &mu)
//	go sumOfPart(data, 500, 750, &totalSum, &wg, &mu)
//	go sumOfPart(data, 750, 1000, &totalSum, &wg, &mu)
//
//	wg.Wait()
//	fmt.Println("total sum =", totalSum)
//}
