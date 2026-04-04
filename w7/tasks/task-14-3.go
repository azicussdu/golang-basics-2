package main

func uniqueWords(words []string) []string {
	mp := map[string]bool{}
	result := []string{}

	for _, word := range words {
		if !mp[word] {
			mp[word] = true
			result = append(result, word)
		}
	}

	return result
}

//func main() {
//	slice := []string{"go", "is", "go", "ok", "bro", "is"}
//	resultSlice := uniqueWords(slice)
//	fmt.Println(resultSlice)
//}
