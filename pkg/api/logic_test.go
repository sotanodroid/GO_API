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

// Test function to get all books
func TestGetBooks(t *testing.T) {
	srv, ctx := setup()

	book, err := srv.GetAllBooks(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	testBook := []models.Book{}

	assert.IsType(t, testBook, book)
}

// Test that book correctly created
func TestCreateBook(t *testing.T) {
	srv, ctx := setup()

	author := models.Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	resp, err := srv.CreateNewBook(ctx, "12345", "Test Book", author)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	assert.Equal(t, resp, "Created")
}

func TestGetBook(t *testing.T) {
	srv, ctx := setup()

	book := models.Book{}
	resp, err := srv.GetBook(ctx, "1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	assert.IsType(t, *resp, book)

}

func TestUpdateBook(t *testing.T) {
	srv, ctx := setup()

	payload := struct {
		Isbn  string
		Title string
	}{
		"123456",
		"Updated Title",
	}

	oldBook, err := srv.GetBook(ctx, "1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	assert.NotEqual(t, oldBook.Isbn, payload.Isbn)
	assert.NotEqual(t, oldBook.Title, payload.Title)

	resp, err := srv.UpdateBook(ctx, strconv.Itoa(oldBook.ID), payload.Isbn, payload.Title)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	book, err := srv.GetBook(ctx, "1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	assert.Equal(t, resp, "Updated")
	assert.Equal(t, book.Isbn, payload.Isbn)
	assert.Equal(t, book.Title, payload.Title)

	// TODO Использовать фикстуры и тестовую БД
	srv.UpdateBook(ctx, strconv.Itoa(oldBook.ID), oldBook.Isbn, oldBook.Title)

}

func setup() (srv Service, ctx context.Context) {
	var logger log.Logger
	var db *pgx.Conn
	ctx = context.Background()

	{
		var err error

		db, err = pgx.Connect(
			ctx,
			"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	repository := models.NewRepo(db, logger)
	srv = NewService(repository, logger)

	return NewService(repository, logger), context.Background()
}
