package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sotanodroid/GO_API/pkg/models"
)

func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	books, err := models.AllBooks()

	if err != nil {
		log.Fatalln("Error in getAllBooks: ", err)
	}

	json.NewEncoder(writer).Encode(books)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "Application/json")
	book := new(models.Book)
	_ = json.NewDecoder(request.Body).Decode(&book)

	if err := models.CreateBook(book); err != nil {
		log.Fatalln("Error in CreateBook: ", err)
	}

	json.NewEncoder(writer).Encode(book)

}

// func getBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(writer).Encode(item)
// 			return
// 		}
// 	}

// 	json.NewEncoder(writer).Encode(&models.Book{})
// }

// func updateBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book models.Book
// 			_ = json.NewDecoder(request.Body).Decode(&book)
// 			book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID
// 			books = append(books, book)
// 			json.NewEncoder(writer).Encode(book)
// 			return
// 		}
// 	}
// }

// func deleteBook(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-type", "Application/json")
// 	params := mux.Vars(request)

// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// }
