package main

import (
	"fmt"
)

type FileSizeExceededError struct{ SizeMB int }

func (fe *FileSizeExceededError) Error() string {
	return fmt.Sprintf("Max 5MB, but your file is %d MB\n", fe.SizeMB)
}

type InvalidFormatError struct{ Format string }

func (fe *InvalidFormatError) Error() string {
	return fmt.Sprintf("Can not upload %s format\n", fe.Format)
}

func UploadFile(filename string, sizeMB int, format string) error {
	if sizeMB > 5 {
		return fmt.Errorf("upload failed: %w\n", &FileSizeExceededError{SizeMB: sizeMB})
	} else if format == ".exe" {
		return fmt.Errorf("upload failed: %w\n", &InvalidFormatError{Format: format})
	}

	fmt.Printf("file: %s is uploaded successfully\n", filename)
	return nil
}

//func main() {
//if err := UploadFile("myfile", 4, ".exe"); err != nil {
//	var fse *FileSizeExceededError
//	var ife *InvalidFormatError
//
//	switch {
//	case errors.As(err, &fse):
//		fmt.Println(fse.Error(), fse.SizeMB)
//	case errors.As(err, &ife):
//		fmt.Println(ife.Error(), ife.Format)
//	}
//
//if errors.As(err, &fse) {
//	fmt.Println(fse.Error(), fse.SizeMB)
//} else if errors.As(err, &ife) {
//	fmt.Println(ife.Error(), ife.Format)
//}
//}

//if err := UploadFile("myfile", 7, ".txt"); err != nil {
//	innerErr := errors.Unwrap(err)
//	if innerErr != nil {
//		fmt.Println(innerErr.Error())
//	}
//}
//}
