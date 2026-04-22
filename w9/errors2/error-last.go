package main

import (
	"fmt"
)

type NoFileError struct {
	Filename string
}

func (f *NoFileError) Error() string {
	return fmt.Sprintf("File not found error, filename: %s\n", f.Filename)
}

//func main() {
//	nfErr := &NoFileError{Filename: "logo.png"}
//	err := fmt.Errorf("Couldn't find a file %w\n", nfErr)
//
//	var e *NoFileError
//	if errors.As(err, &e) {
//		fmt.Println("It is an error: ", e.Filename)
//	}
//}
