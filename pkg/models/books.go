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
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

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
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		query,
		book.Isbn,
		book.Title,
		book.Author.ID,
	)

	var id uint64
	err = row.Scan(&id)

	if err != nil {
		return errors.New("Error inserting data into Books")
	}

	return nil
}
