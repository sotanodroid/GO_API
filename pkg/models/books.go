package models

import (
	"context"
	"errors"
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

// TODO
// 	getBook() (*Book, error)
// 	createBook(book *Book) error
// 	updateBook(book *Book) error
// 	deleteBook(book *Book) error

//AllBooks Select all books from db
func AllBooks() ([]*Book, error) {
	query := "SELECT * FROM goapi.books;"

	rows, err := db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		if err := rows.Scan(
			&bk.ID,
			&bk.Isbn,
			&bk.Title,
			&bk.Author.ID,
		); err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

//CreateBook creates a new book
func CreateBook(book *Book) error {
	query := `
		INSERT INTO goapi.books 
		(
			isbn,
			title,
			author
		) 
		VALUES
		(
			$1,
			$2,
			$3
		)
		RETURNING id;`

	_, err := db.Exec(
		context.Background(),
		query,
		book.Isbn,
		book.Title,
		book.Author.ID,
	)

	if err != nil {
		return errors.New("Error inserting data into Books")
	}

	return nil
}
