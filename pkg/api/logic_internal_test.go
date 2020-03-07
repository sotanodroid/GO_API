package api

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jackc/pgx/v4"
	"github.com/sotanodroid/GO_API/pkg/models"
)

func TestRepository(t *testing.T) {
	srv, ctx := setup()

	var oldBook models.Book

	payload := struct {
		Isbn  string
		Title string
	}{
		"123456",
		"Updated Title",
	}

	author := models.Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	{
		resp, err := srv.CreateNewBook(ctx, "12345", "Test Book", author)
		if err != nil {
			t.Errorf("Error CreateNewBook: %s", err)
		}

		assert.Equal(t, resp, "Created")
	}

	{
		allBooks, err := srv.GetAllBooks(ctx)
		if err != nil {
			t.Errorf("Error GetAllBooks: %s", err)
		}

		oldBook = allBooks[len(allBooks)-1]

		assert.NotEqual(t, oldBook.Isbn, payload.Isbn)
		assert.NotEqual(t, oldBook.Title, payload.Title)
	}

	{
		resp, err := srv.UpdateBook(
			ctx,
			strconv.Itoa(oldBook.ID),
			payload.Isbn,
			payload.Title,
		)
		if err != nil {
			t.Errorf("Error UpdateBook: %s", err)
		}

		book, err := srv.GetBook(ctx, strconv.Itoa(oldBook.ID))
		if err != nil {
			t.Errorf("Error GetBook: %s", err)
		}

		assert.Equal(t, resp, "Updated")
		assert.Equal(t, book.Isbn, payload.Isbn)
		assert.Equal(t, book.Title, payload.Title)
	}

	{
		resp, err := srv.DeleteBook(ctx, strconv.Itoa(oldBook.ID))
		if err != nil {
			t.Errorf("Error DeleteBook: %s", err)
		}

		assert.Equal(t, resp, "Deleted")
	}
}

func setup() (srv Service, ctx context.Context) {
	var logger log.Logger
	var db *pgx.Conn
	ctx = context.Background()

	{
		var err error

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			dbURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		}

		db, err = pgx.Connect(
			ctx,
			dbURL,
		)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	repository := models.NewRepo(db, logger)

	return NewService(repository, logger), context.Background()
}
