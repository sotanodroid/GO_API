package api

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/sotanodroid/GO_API/pkg/models"
)

// Service of microservice
type service struct {
	repository models.Repository
	logger     log.Logger
}

// NewService returns new instance of servise
func NewService(rep models.Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// GetAllBooks gets all books
func (s service) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	logger := log.With(s.logger, "method", "getAllBooks")

	books, err := s.repository.AllBooks(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return books, nil
}

// CreateNewBook creates new book
func (s service) CreateNewBook(
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
func (s service) GetBook(ctx context.Context, id string) (*models.Book, error) {
	logger := log.With(s.logger, "method", "getBook")
	book, err := s.repository.GetBook(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return book, nil
}

// func getBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	book, err := models.GetBook(params["id"])

// 	if err != nil {
// 		level.Error(logger).Log("err", err)
// 	}

// 	json.NewEncoder(writer).Encode(book)
// }

// func updateBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)
// 	book := new(models.Book)

// 	_ = json.NewDecoder(request.Body).Decode(&book)
// 	book.ID, _ = strconv.Atoi(params["id"])

// 	if err := models.UpdateBook(book); err != nil {
// 		level.Error(logger).Log("err", err)
// 	}

// 	json.NewEncoder(writer).Encode(book)
// }

// func deleteBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	if err := models.DeleteBook(params["id"]); err != nil {
// 		level.Error(logger).Log("err", err)
// 	}

// 	writer.Write([]byte("Deleted"))
// }
