package main

import (
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Split(strings.ToLower(text), " ") // words = ["go", "is" ...]
	myMap := map[string]int{}

	for _, word := range words {
		myMap[word]++
	}

	return myMap
}

//func main() {
//	str := "Go is the best lang in the world go do go"
//	resultMap := wordFrequency(str)
//	fmt.Println(resultMap)
//}
