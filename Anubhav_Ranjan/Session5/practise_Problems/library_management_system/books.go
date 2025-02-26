package library_management_system

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func AddBook(db *[]Book, title string, author string, isbn string) {
	newBook := Book{Title: title, Author: author, ISBN: isbn}
	*db = append(*db, newBook)
	fmt.Printf(" Book Added: %+v\n", newBook)
}

func RemoveBook(db *[]Book, isbn string) error {
	for i, book := range *db {
		if book.ISBN == isbn {
			*db = append((*db)[:i], (*db)[i+1:]...)
			fmt.Printf(" Book with ISBN %s is Removed..\n", isbn)
			return nil
		}
	}

	return errors.New("Book NOT Found")
}

func SearchBookByTitle(db *[]Book, title string) (Book, error) {
	for _, book := range *db {
		if book.Title == title {
			return book, nil
		}
	}
	return Book{}, errors.New("Book NOT Found")
}

func DisplayBooks(db *[]Book) {
	if len(*db) == 0 {
		fmt.Println("No Books Found")
		return
	}

	fmt.Println("\n Book Records:")
	for _, book := range *db {
		fmt.Printf("Title : %s | Author : %s | ISBN : %s\n", book.Title, book.Author, book.ISBN)
	}
}
