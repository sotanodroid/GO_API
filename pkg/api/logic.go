package api

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/sotanodroid/GO_API/pkg/models"
)

type app struct {
	repository models.Repository
	logger     log.Logger
}

// NewService returns new instance of servise
func NewService(rep models.Repository, logger log.Logger) Service {
	return &app{
		repository: rep,
		logger:     logger,
	}
}

// GetAllBooks gets all books
func (s app) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	logger := log.With(s.logger, "method", "getAllBooks")

	books, err := s.repository.AllBooks(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return books, nil
}

// CreateNewBook creates new book
func (s app) CreateNewBook(
	ctx context.Context,
	isbn, title string,
	author models.Author,
) (string, error) {
	logger := log.With(s.logger, "method", "createBook")
	book := models.Book{
		Isbn:   isbn,
		Title:  title,
		Author: author,
	}

	if err := s.repository.CreateBook(ctx, book); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	return "Created", nil
}

// GetBook gets single book
func (s app) GetBook(ctx context.Context, id string) (*models.Book, error) {
	logger := log.With(s.logger, "method", "getBook")
	book, err := s.repository.GetBook(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return book, nil
}

func (s app) UpdateBook(ctx context.Context, id, Isbn, Title string) (string, error) {
	logger := log.With(s.logger, "method", "getBook")

	if err := s.repository.UpdateBook(ctx, id, Isbn, Title); err != nil {
		level.Error(logger).Log("err", err)
	}

	return "Updated", nil
}

// func deleteBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	if err := models.DeleteBook(params["id"]); err != nil {
// 		level.Error(logger).Log("err", err)
// 	}

// 	writer.Write([]byte("Deleted"))
// }
