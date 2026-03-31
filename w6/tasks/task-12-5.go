package main

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Name  string
	Books []Book
}

func NewLibrary(name string) *Library {
	return &Library{
		Name:  name,
		Books: []Book{},
	}
}

func (lib *Library) AddBook(b Book) {
	lib.Books = append(lib.Books, b)
}

func (lib *Library) FindByTitle(title string) *Book {
	for i := range lib.Books {
		if lib.Books[i].Title == title {
			return &lib.Books[i]
		}
	}
	return nil
}

func (lib *Library) BooksByAuthor(author string) []Book {
	var books []Book

	for i := range lib.Books {
		if lib.Books[i].Author == author {
			books = append(books, lib.Books[i])
		}
	}

	return books
}

func (lib *Library) RemoveBook(title string) bool {
	// [b1 b2 b3 b4 b5 b6]  i=2
	for i := range lib.Books {
		if lib.Books[i].Title == title {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			return true
		}
	}
	return false
}

//func main() {
//	library := NewLibrary("im. Abaya")
//	fmt.Println(library)
//
//	library.AddBook(Book{Title: "Book1", Author: "Nurbakhyt", Year: 2000})
//	library.AddBook(Book{Title: "Book2", Author: "Nurbakhyt", Year: 2000})
//	library.AddBook(Book{Title: "Book3", Author: "Ualikhan", Year: 2000})
//	library.AddBook(Book{Title: "Book4", Author: "Ualikhan", Year: 2000})
//
//	fmt.Println(library)
//
//	foundBook := library.FindByTitle("Book7")
//	fmt.Println(foundBook)
//
//	booksByAuthor := library.BooksByAuthor("Ualikhan")
//	fmt.Println(booksByAuthor)
//
//	removed := library.RemoveBook("Book3")
//	if removed == true {
//		fmt.Println("Oshirdik")
//	}
//
//	fmt.Println(library)
//}
