package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func getUser(id int) error {
	if id != 1 {
		return fmt.Errorf("User with id: %d not found (%w)\n", id, ErrNotFound)
	}
	return nil
}

//func main() {
//	err := getUser(10)
//
//	if errors.Is(err, ErrNotFound) {
//		fmt.Println("User is not found")
//	}
//}
