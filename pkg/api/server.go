package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RunServer handles main routing for server
func RunServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", getAllBooks).Methods("GET")
	// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}