package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"github.com/sotanodroid/GO_API/pkg/models"
)

type (
	//GetBooksRequest struct
	GetBooksRequest struct{}

	//GetBooksResponse struct of slice of Books
	GetBooksResponse struct {
		Books []models.Book `json:"books"`
	}

	//CreateBookRequest Isbn, Title, Author struct
	CreateBookRequest struct {
		Isbn   string        `json:"isbn"`
		Title  string        `json:"title"`
		Author models.Author `json:"author"`
	}

	//OkResponse struct
	OkResponse struct {
		Ok string `json:"ok"`
	}

	// IDRequest gets single book
	IDRequest struct {
		ID string
	}

	// GetBookResponse returns single book
	GetBookResponse struct {
		ID     int           `json:"id"`
		Isbn   string        `json:"isbn"`
		Title  string        `json:"title"`
		Author models.Author `json:"author"`
	}

	// UpdateBookRequest updates book
	UpdateBookRequest struct {
		ID    string `json:"id"`
		Isbn  string `json:"isbn"`
		Title string `json:"title"`
	}
)

//Endpoints holds endpoints
type Endpoints struct {
	GetBooks   endpoint.Endpoint
	CreateBook endpoint.Endpoint
	GetBook    endpoint.Endpoint
	UpdateBook endpoint.Endpoint
	DeleteBook endpoint.Endpoint
}

//MakeEndpoints makes endpoints to handle requests
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetBooks:   makeGetBooksEndpoint(s),
		CreateBook: makeCreateBooksEndpoints(s),
		GetBook:    makeGetBookEndpoint(s),
		UpdateBook: makeUpdateBookEndpoint(s),
		DeleteBook: makeDeleteBookEndpoint(s),
	}
}

func makeGetBooksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		books, err := s.GetAllBooks(ctx)

		return GetBooksResponse{Books: books}, err
	}
}

func makeCreateBooksEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBookRequest)
		author := models.Author{
			Firstname: req.Author.Firstname,
			Lastname:  req.Author.Lastname,
		}
		ok, err := s.CreateNewBook(ctx, req.Isbn, req.Title, author)

		return OkResponse{Ok: ok}, err
	}
}

func makeGetBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IDRequest)
		book, err := s.GetBook(ctx, req.ID)

		response := GetBookResponse{
			ID:     book.ID,
			Isbn:   book.Isbn,
			Title:  book.Title,
			Author: book.Author,
		}

		return response, err
	}
}

func makeUpdateBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBookRequest)
		ok, err := s.UpdateBook(ctx, req.ID, req.Isbn, req.Title)

		response := OkResponse{Ok: ok}

		return response, err
	}
}

func makeDeleteBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IDRequest)
		ok, err := s.DeleteBook(ctx, req.ID)

		response := OkResponse{Ok: ok}

		return response, err
	}
}

func encodeResponse(ctx context.Context, writter http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writter).Encode(response)
}

func decodeBookRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateBookRequest

	if request.Body != nil {
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}
		return req, nil
	}

	return nil, nil
}

func decodeIDRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)

	return IDRequest{ID: params["id"]}, nil
}

func decodePutRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)

	var req UpdateBookRequest
	if request.Body != nil {
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			return nil, err
		}
		req.ID = params["id"]

		return req, nil
	}

	return nil, nil
}
