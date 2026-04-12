package main

import "fmt"

type FileStorage interface {
	Open(filename string) bool
	Read() (string, bool)
	Write(content string) bool
	Close() bool
}

type MockFileStorage struct {
	isOpen   bool
	filename string
	content  string
}

func (mfs *MockFileStorage) Close() bool {
	mfs.isOpen = false
	return true
}

func (mfs *MockFileStorage) Read() (string, bool) {
	if !mfs.isOpen {
		return "", false
	}
	return mfs.content, true
}

func (mfs *MockFileStorage) Write(content string) bool {
	if !mfs.isOpen {
		return false
	}
	mfs.content += content
	return true
}

func (mfs *MockFileStorage) Open(filename string) bool {
	mfs.isOpen = true
	mfs.filename = filename
	mfs.content = ""
	return true
}

func TestStorage(fs FileStorage, filename string) bool {
	if fs.Open(filename) {
		written := fs.Write("Salam aleikum jigitter!!!")
		if !written {
			return false
		}
		msg, ok := fs.Read()
		if ok {
			fmt.Println(msg)
		} else {
			return false
		}
		fs.Close()
		return true
	} else {
		return false
	}
}

//func main() {
//	m := &MockFileStorage{}
//	ok := TestStorage(m, "myfile.txt")
//
//	if ok {
//		fmt.Println("Test is done successfully")
//	} else {
//		fmt.Println("Test is not correct")
//	}
//}
