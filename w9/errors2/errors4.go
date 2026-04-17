package main

import (
	"fmt"
)

type FileNotFound struct {
	Filename string
}

func (f *FileNotFound) Error() string {
	return fmt.Sprintf("File not found error, filename: %s\n", f.Filename)
}

func findFile(fname string) error {
	if fname != "myfile.txt" {
		return &FileNotFound{Filename: fname}
	}
	return nil
}

func load() error {
	err := findFile("logo.png")
	if err != nil {
		return fmt.Errorf("no file: %w", err)
	}
	return nil
}

//func main() {
//	err := load()
//	var fNotFound *FileNotFound
//	if errors.As(err, &fNotFound) {
//		fmt.Println("Path:", fNotFound.Filename)
//	}
//}
