package models

import (
	"context"
)

// Book model Struct
type Book struct {
	ID     int    `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author model Struct
type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Repository is an interface to a book structure
type Repository interface {
	AllBooks(ctx context.Context) ([]Book, error)
	CreateBook(ctx context.Context, book Book) error
}
