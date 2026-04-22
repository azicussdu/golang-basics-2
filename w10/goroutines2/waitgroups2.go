package main

import (
	"fmt"
	"sync"
)

func correctExample() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("Worker %d\n", id)
		}(i, &wg)
	}

	wg.Wait()
}

//func main() {
//	correctExample()
//}
