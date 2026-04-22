package main

import (
	"errors"
	"fmt"
)

var ErrInvalidToken = errors.New("invalid token")

func Authenticate(token string) error {
	if token != "almaty" {
		return fmt.Errorf("auth failed: %w\n", ErrInvalidToken)
	}
	// tut otti
	return nil
}

//func main() {
//	if err := Authenticate("almaty"); err != nil {
//		innerErr := errors.Unwrap(err)
//		if innerErr != nil {
//			fmt.Println("Wrapped error: ", innerErr.Error())
//		}
//	} else {
//		fmt.Println("Success")
//	}
//}
