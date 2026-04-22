package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value++
}

func (sc *SafeCounter) GetValue() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.value
}

func (sc *SafeCounter) Reset() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value = 0
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	fmt.Printf("Значение счетчика: %d\n", counter.GetValue())
	counter.Reset() // Сбрасываем счетчик
	fmt.Printf("После сброса: %d\n", counter.GetValue())
}
