package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
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

	//CreateBookResponse struct
	CreateBookResponse struct {
		Ok string `json: ok`
	}
)

//Endpoints holds endpoints
type Endpoints struct {
	GetBooks   endpoint.Endpoint
	CreateBook endpoint.Endpoint
}

//MakeEndpoints makes endpoints to handle requests
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetBooks:   makeGetBooksEndpoint(s),
		CreateBook: makeCreateBooksEndpoints(s),
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
		return CreateBookResponse{Ok: ok}, err
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
