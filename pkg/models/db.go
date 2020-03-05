package models

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/jackc/pgx/v4"
)

type repo struct {
	db     *pgx.Conn
	logger log.Logger
}

// NewRepo creates new repo
func NewRepo(db *pgx.Conn, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

//AllBooks Select all books from db
func (r *repo) AllBooks(ctx context.Context) ([]Book, error) {
	const query = `
		SELECT b.id, b.Isbn, b.Title, a.id, a.firstname, a.lastname
		FROM goapi.books as b
		JOIN goapi.authors as a
		ON b.author = a.id;`

	rows, err := r.db.Query(context.Background(), query)

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
func (r *repo) CreateBook(ctx context.Context, book Book) error {
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

	commandTag, err := r.db.Exec(
		context.Background(),
		query,
		book.Isbn,
		book.Title,
		book.Author.Firstname,
		book.Author.Lastname,
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
func (r *repo) GetBook(ctx context.Context, id string) (*Book, error) {
	const query = `
		SELECT b.id, b.Isbn, b.Title, a.id, a.firstname, a.lastname
		FROM goapi.books as b
		JOIN goapi.authors as a
		ON b.author = a.id
		WHERE b.id = $1;`

	row := r.db.QueryRow(context.Background(), query, id)

	var bk Book

	if err := row.Scan(
		&bk.ID,
		&bk.Isbn,
		&bk.Title,
		&bk.Author.ID,
		&bk.Author.Firstname,
		&bk.Author.Lastname,
	); err != nil {
		return nil, err
	}

	return &bk, nil
}

// UpdateBook updates book by it's ID
func (r *repo) UpdateBook(ctx context.Context, id, Isbn, Title string) error {
	const query = `
		UPDATE goapi.books
		SET isbn = $2, title = $3
		WHERE
		id = $1;`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
		Isbn,
		Title,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteBook would delete book by id
func (r *repo) DeleteBook(ctx context.Context, id string) error {
	const query = `
		DELETE FROM goapi.books
		WHERE id = $1;`

	_, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}
