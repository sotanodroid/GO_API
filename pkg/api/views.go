package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sotanodroid/GO_API/pkg/db"
)

var books []db.Book

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

	json.NewEncoder(writer).Encode(&db.Book{})
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	var book db.Book
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
			var book db.Book
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
