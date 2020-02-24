package api

import (
	"context"

	"github.com/sotanodroid/GO_API/pkg/models"
)

// Service is a main microservice interface
type Service interface {
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	CreateNewBook(ctx context.Context, isbn, title string, author models.Author) (string, error)
}