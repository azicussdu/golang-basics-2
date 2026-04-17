package main

func mayPanic() {
	panic("Критическая ошибка!")
}

//func main() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Восстановление после паники:", r)
//		}
//	}()
//	mayPanic()
//	fmt.Println("Если будет паника то это сообщение не увидим")
//}
