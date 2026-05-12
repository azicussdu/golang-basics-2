package main

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

//func main() {
//	sum := applyOperation(5, 3, add)          // функцию отправляем как параметр в другую функцию
//	product := applyOperation(5, 3, multiply) // тут тоже самое
//	fmt.Println("Сумма:", sum)                // Вывод: Sum: 8
//	fmt.Println("Product:", product)          // Вывод: Product: 15
//}
