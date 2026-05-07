package main

import (
	"fmt"
	"time"
)

// Worker pool
func worker(id int, jobs chan int) {
	for job := range jobs {
		fmt.Printf("Воркер %d обрабатывает задачу %d\n", id, job)
		time.Sleep(time.Second)
	}
}

//func main() {
//	jobs := make(chan int, 5)
//	for i := 1; i <= 3; i++ {
//		go worker(i, jobs)
//	}
//	for j := 1; j <= 10; j++ {
//		time.Sleep(10 * time.Millisecond)
//		jobs <- j
//		fmt.Println("Отправлена задача", j)
//	}
//	close(jobs)
//	time.Sleep(5 * time.Second)
//}
