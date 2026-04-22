package main

import (
	"errors"
	"fmt"
)

type FileStorage interface {
	Open(filename string) error
	Read() (string, error)
	Write(content string) error
	Close() error
}

type MockFileStorage struct {
	isOpen   bool
	filename string
	content  string
}

func (mfs *MockFileStorage) Close() error {
	if !mfs.isOpen { // uje jabyk bolsa
		return errors.New("file is already closed")
	}
	mfs.isOpen = false
	return nil
}

func (mfs *MockFileStorage) Read() (string, error) {
	if !mfs.isOpen {
		return "", errors.New("can not read when file is closed")
	}
	return mfs.content, nil
}

func (mfs *MockFileStorage) Write(content string) error {
	if !mfs.isOpen {
		return errors.New("can not write when file is closed")
	}
	mfs.content += content
	return nil
}

func (mfs *MockFileStorage) Open(filename string) error {
	if len(filename) == 0 {
		return errors.New("file name must contain at least one character")
	} else if mfs.isOpen {
		return errors.New("already open")
	}
	mfs.isOpen = true
	mfs.filename = filename
	mfs.content = ""
	return nil
}

func TestStorage(fs FileStorage, filename string) error {
	err := fs.Open(filename)
	if err != nil {
		return err
	}

	if err = fs.Write("Salam aleikum jigitter!!!"); err != nil {
		return err
	}

	var msg string
	if msg, err = fs.Read(); err != nil {
		return err
	}
	fmt.Println(msg)

	if err = fs.Close(); err != nil {
		return err
	}

	return nil
}

//func main() {
//	m := &MockFileStorage{}
//	err := TestStorage(m, "myfile.txt")
//
//	if err == nil {
//		fmt.Println("Test is done successfully")
//	} else {
//		fmt.Println("Test is not correct")
//		fmt.Println(err.Error())
//	}
//}
