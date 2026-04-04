package main

import (
	"strings"
)

func charCount(s string) map[rune]int {
	s = strings.ToLower(s)
	myMap := make(map[rune]int)

	for _, letter := range s {
		myMap[letter]++
	}

	return myMap
}

//func main() {
//	str := "Successes"
//	resultMap := charCount(str)
//	for k, v := range resultMap {
//		fmt.Printf("%c - %d\n", k, v)
//	}
//}
