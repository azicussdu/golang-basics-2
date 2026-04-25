package main

import (
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, slice []uint) {
	defer wg.Done()

	res := uint(1)
	for i := 2; i <= n; i++ {
		res *= uint(i)
	}
	slice[n-1] = res
}

//func main() {
//	var wg sync.WaitGroup
//	result := make([]uint, 14)
//
//	for i := 1; i <= 14; i++ {
//		wg.Add(1)
//		go factorial(i, &wg, result)
//	}
//
//	wg.Wait()
//
//	fmt.Println(result)
//}
