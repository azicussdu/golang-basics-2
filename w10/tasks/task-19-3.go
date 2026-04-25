package main

import (
	"fmt"
	"sync"
)

func averageOfRange(data []int, startIndex, endIndex, index int, averages []float32, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for i := startIndex; i < endIndex; i++ {
		sum += data[i]
	}

	averages[index] = float32(sum) / float32(endIndex-startIndex)

	fmt.Printf("Gourotine #%d [%d:%d] avg=%.2f\n", index+1, startIndex, endIndex, averages[index])
}

//func main() {
//	var wg sync.WaitGroup
//
//	totalNums := 1000
//	gCount := 5
//	numPerGr := totalNums / gCount // = 100
//
//	data := make([]int, totalNums)
//	for i := 0; i < totalNums; i++ {
//		data[i] = rand.IntN(1000)
//	}
//
//	averages := make([]float32, gCount)
//
//	for i := 0; i < gCount; i++ {
//		wg.Add(1)
//		startIndex := i * numPerGr        // 0		100
//		endIndex := startIndex + numPerGr // 100	200
//
//		go averageOfRange(data, startIndex, endIndex, i, averages, &wg)
//	}
//
//	wg.Wait()
//
//	finalAvg := float32(0.0)
//	for i := 0; i < gCount; i++ {
//		finalAvg += averages[i]
//	}
//	fmt.Println("final average:", finalAvg/float32(gCount))
//}
