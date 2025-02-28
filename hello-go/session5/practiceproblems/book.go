package main

import (
	"fmt"
	"strings"
)


type Book struct {
	Title  string
	Author string
	ISBN   string
}


func searchBookByTitle(books []Book, title string) *Book {
	for _, book := range books {
		if strings.EqualFold(book.Title, title) { 
			return &book
		}
	}
	return nil
}

func main() {
	
	books := []Book{
		{"The Great Gatsby", "F. Scott Fitzgerald", "9780743273565"},
		{"To Kill a Mockingbird", "Harper Lee", "9780061120084"},
		{"1984", "George Orwell", "9780451524935"},
	}


	var title string
	fmt.Print("Enter the book title to search: ")
	fmt.Scanln(&title)

	
	foundBook := searchBookByTitle(books, title)

	
	if foundBook != nil {
		fmt.Printf("Book Found:\nTitle: %s\nAuthor: %s\nISBN: %s\n", foundBook.Title, foundBook.Author, foundBook.ISBN)
	} else {
		fmt.Println("Book not found.")
	}
}
