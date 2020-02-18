package models

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

//InitDB initialize db connection
func InitDB(dataSourceName string) {
	var err error
	ctx := context.Background()

	db, err = pgx.Connect(ctx, dataSourceName)
	if err != nil {
		log.Println("Error connecting to database", err)
	}

}

// TODO
// 	createBook(book *Book) error
// 	updateBook(book *Book) error
// 	deleteBook(book *Book) error

//AllBooks Select all books from db
func AllBooks() ([]Book, error) {
	const query = `
		SELECT b.id, b.Isbn, b.Title, a.id, a.firstname, a.lastname
		FROM goapi.books as b
		JOIN goapi.authors as a
		ON b.author = a.id;
		`
	rows, err := db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := new(Book)
		if err := rows.Scan(
			&bk.ID,
			&bk.Isbn,
			&bk.Title,
			&bk.Author.ID,
			&bk.Author.Firstname,
			&bk.Author.Lastname,
		); err != nil {
			return nil, err
		}
		bks = append(bks, *bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

//CreateBook creates a new book
func CreateBook(book *Book) error {
	const query = `
		INSERT INTO goapi.books 
		(isbn, title, author) 
		VALUES
		(
			$1,
			$2,
			(
				SELECT id
				FROM goapi.authors
				WHERE firstname = $3
				AND lastname = $4
				LIMIT 1
			)
		);`

	commandTag, err := db.Exec(
		context.Background(),
		query,
		&book.Isbn,
		&book.Title,
		&book.Author.Firstname,
		&book.Author.Lastname,
	)

	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("Error Executing INSERT on CreateBook")
	}
	return nil
}

//GetBook gets single book
func GetBook(id string) (*Book, error) {
	const query = `
		SELECT b.id, b.Isbn, b.Title, a.id, a.firstname, a.lastname
		FROM goapi.books as b
		JOIN goapi.authors as a
		ON b.author = a.id
		WHERE b.id = $1;
		`
	rows, err := db.Query(context.Background(), query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bk := new(Book)
	for rows.Next() {
		if err := rows.Scan(
			&bk.ID,
			&bk.Isbn,
			&bk.Title,
			&bk.Author.ID,
			&bk.Author.Firstname,
			&bk.Author.Lastname,
		); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bk, nil
}
