package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader // kak-budto "Read() string" jaza salgandai
	Writer // kak-budto "Write(data string)" jaza salgandai
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

func PrintContent(rw ReadWriter) {
	rw.Write("Hello, world!")
	fmt.Println("Reading:", rw.Read())
}

//func main() {
//	file := &File{}
//	PrintContent(file) // Reading: Hello, world!
//}
