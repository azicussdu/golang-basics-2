package main

import "fmt"

type Book struct {
	ID     string
	Title  string
	Author string
}

type Member struct {
	ID   string
	Name string
}

type Library struct {
	BooksMap      map[string]Book
	MembersMap    map[string]Member
	TakenBooksMap map[string]string
}

func NewLibrary() *Library {
	lib := Library{
		BooksMap:      make(map[string]Book),
		MembersMap:    make(map[string]Member),
		TakenBooksMap: make(map[string]string),
	}
	return &lib
}

func (lib *Library) AddBook(b Book) {
	lib.BooksMap[b.ID] = b
}

func (lib *Library) AddMember(m Member) {
	lib.MembersMap[m.ID] = m
}

func (lib *Library) BorrowBook(bookID, memberID string) bool {
	_, bookExists := lib.BooksMap[bookID]
	_, memberExists := lib.MembersMap[memberID]

	if !bookExists || !memberExists {
		return false
	}

	if _, takenBook := lib.TakenBooksMap[bookID]; takenBook {
		return false
	}

	lib.TakenBooksMap[bookID] = memberID
	return true
}

func (lib *Library) ReturnBook(bookID string) bool {
	if _, takenBook := lib.TakenBooksMap[bookID]; takenBook == false {
		return false
	}

	delete(lib.TakenBooksMap, bookID)
	return true
}

func (lib *Library) ListBorrowedBooks() {
	if len(lib.TakenBooksMap) == 0 {
		fmt.Println("Nobody took no book")
		return
	}

	for bookID, memberID := range lib.TakenBooksMap {
		book := lib.BooksMap[bookID]
		member := lib.MembersMap[memberID]
		fmt.Printf("%s borrowed book %s\n", member.Name, book.Title)
	}
}

//func main() {
//
//	library := NewLibrary()
//
//	b1 := Book{ID: "b1", Title: "Book1", Author: "Arman"}
//	b2 := Book{ID: "b2", Title: "Book2", Author: "Daulet"}
//	library.AddBook(b1)
//	library.AddBook(b2)
//
//	m1 := Member{ID: "m1", Name: "Beka"}
//	m2 := Member{ID: "m2", Name: "Seka"}
//	library.AddMember(m1)
//	library.AddMember(m2)
//
//	library.ListBorrowedBooks()
//	library.BorrowBook(b1.ID, m1.ID)
//	library.BorrowBook(b2.ID, m1.ID)
//
//	library.ListBorrowedBooks()
//
//	library.ReturnBook(b2.ID)
//	fmt.Println("Returned book b2")
//
//	library.ListBorrowedBooks()
//}
