package main

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

//func main() {
//	double := makeMultiplier(2) // мы вызываем функцию, которая возвращает другую функцию
//	triple := makeMultiplier(3)
//	fmt.Println(double(5)) // Вывод: 10
//	fmt.Println(triple(5)) // Вывод: 15
//}
