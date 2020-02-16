package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sotanodroid/GO_API/pkg/db"
)

// RunServer handles main routing for server
func RunServer() {
	// Init router
	router := mux.NewRouter()

	// Mock data
	books = append(books, db.Book{ID: "1", Isbn: "153223", Title: "Book One",
		Author: &db.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, db.Book{ID: "2", Isbn: "153235", Title: "Book Two",
		Author: &db.Author{Firstname: "Steve", Lastname: "Smith"}})

	//Route handlers
	router.HandleFunc("/api/books", getAllBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
