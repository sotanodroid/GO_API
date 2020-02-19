package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sotanodroid/GO_API/pkg/models"
)

func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	books, err := models.AllBooks()

	if err != nil {
		log.Println("Error in getAllBooks: ", err)
	}

	json.NewEncoder(writer).Encode(books)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	book := new(models.Book)
	_ = json.NewDecoder(request.Body).Decode(&book)

	if err := models.CreateBook(book); err != nil {
		log.Println("Error in CreateBook: ", err)
	}

	json.NewEncoder(writer).Encode(book)

}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)

	book, err := models.GetBook(params["id"])

	if err != nil {
		log.Println("Error in GetBook: ", err)
	}

	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)
	book := new(models.Book)

	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID, _ = strconv.Atoi(params["id"])

	if err := models.UpdateBook(book); err != nil {
		log.Println("Error in updateBook: ", err)
	}

	json.NewEncoder(writer).Encode(book)
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	params := mux.Vars(request)

	if err := models.DeleteBook(params["id"]); err != nil {
		log.Println("Error in GetBook: ", err)
	}

	writer.Write([]byte("Deleted"))
}
