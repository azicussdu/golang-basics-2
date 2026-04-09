package main

import (
	"fmt"
)

func printAny(a interface{}) {
	fmt.Print("mani: ", a)
	fmt.Printf(" tipi: %T\n", a)
}

func PrintType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Println("unknown type")
	}
}

//func main() {
//var x interface{} // interface{} - тип (int, float64, string, bool)
//
//x = 23
//fmt.Println(x)
//
//x = "Salem"
//fmt.Println(x)
//
//x = 5.6
//fmt.Println(x)
//
//x = false
//fmt.Println(x)

//printAny(40)
//printAny("Salem")
//printAny(true)
//printAny(5.4)

//var x interface{} = "ASds"
//
//y, ok := x.(int)
//if ok {
//	fmt.Println(y + 10)
//}

//PrintType(10)
//PrintType("adasds")
//PrintType(2.3)
//PrintType(true)

//jsonData := `{"name":"Alice","age":25,
//			"hobbies": ["coding", "hiking", "reading"],
//			"active":true}`
//
//var result map[string]interface{}
//err := json.Unmarshal([]byte(jsonData), &result)
//
//if err != nil {
//	return
//}
//
//fmt.Println("Result:", result["age"])
//
//myAge, ok := result["age"].(float64)
//if ok {
//	fmt.Println("cherez 10 let mne budet ", myAge+10, " let")
//}
//}
