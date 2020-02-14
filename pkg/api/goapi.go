package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

// RunServer handles main routing for server
func RunServer() {
	// Init router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "153223", Title: "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "153235", Title: "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	//Route handlers
	router.HandleFunc("/api/books", getAllBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
