package api

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// router.HandleFunc("/api/books", getAllBooks).Methods("GET")
// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// router.HandleFunc("/api/books", createBook).Methods("POST")
// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

// NewHTTPServer creates new server to serve endpoints
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleWare)

	r.Methods("POST").Path("/api/books").Handler(httptransport.NewServer(
		endpoints.CreateBook,
		decodeBookRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/api/books").Handler(httptransport.NewServer(
		endpoints.GetBooks,
		decodeBookRequest,
		encodeResponse,
	))

	// For methods with /{id} need to create unique decode function
	return r
}

func commonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-type", "Application/json")
		next.ServeHTTP(writer, request)
	})
}
