package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) bool
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]*models.Book
	members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		books: map[int]*models.Book{},
		members: map[int]*models.Member{
			1: {
				ID:            1,
				BorrowedBooks: []models.Book{},
				Name:          "Beka",
			},
		},
	}
}

func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = &book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)

}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exist := l.books[bookID]
	if !exist {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exist := l.members[memberID]
	if !exist {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, *book)
	fmt.Println(member.BorrowedBooks)

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	var updatedBorrowedBooks []models.Book
	for _, b := range member.BorrowedBooks {
		if b.ID != bookID {
			updatedBorrowedBooks = append(updatedBorrowedBooks, b)
		}
	}

	member.BorrowedBooks = updatedBorrowedBooks

	book.Status = "Available"

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, *book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
