package main

//func main() {
//	file, err := os.Create("test.txt") // создаем файл
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()                         // Соединение с файлом закроется после как закончиться функция main
//	_, err = file.WriteString("Hello, World!") // пишем в файл
//	if err != nil {
//		fmt.Println(err)
//		return // даже если выйдем тут, то file.Close() выполниться
//	}
//	fmt.Println("Data written to file")
//}
