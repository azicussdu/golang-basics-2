package main

import (
	"fmt"
	"time"
)

func printNumbers(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("GR #%d: (i: %d)\n", id, i)
		time.Sleep(50 * time.Millisecond)
	}
}

//func main() {
//	for i := 1; i <= 3; i++ {
//		go printNumbers(i)
//	}
//	time.Sleep(1 * time.Second)
//}
