package main

func hasDuplicates(nums []int) bool {
	mp := make(map[int]bool)

	for _, num := range nums {
		if mp[num] { // eger num mapada bar bolsa: if mp[num] == true {...}
			return true
		}
		mp[num] = true
	}
	return false
}

//func main() {
//	numbers := []int{2, 3, 4, 5, 7}
//	fmt.Println(hasDuplicates(numbers))
//}
