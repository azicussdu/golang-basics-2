package main

import (
	"fmt"
)

type Kate struct {
	message string
	code    int
}

func (k *Kate) Error() string {
	return fmt.Sprintf("message: %s | error code: %d\n", k.message, k.code)
}

func divide(a, b float64) (float64, error) { // *Kate   Kate
	if b == 0 {
		err := &Kate{message: "Nolge boluge bolmaidy", code: 404}
		return 0, err
	}

	if a == 0 {
		err := &Kate{message: "Noldi sanga bolume bolmaidy", code: 500}
		return 0, err
	}

	res := a / b
	return res, nil
}

//func main() {
//	result, err := divide(23, 0)
//	if err != nil {
//		fmt.Println("Kate bolyp kaldy:", err.Error())
//		return
//	}
//
//	fmt.Println("result =", result)
//}
