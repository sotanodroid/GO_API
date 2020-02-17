package api

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RunServer handles main routing for server
func RunServer() {
	// переделать на джин. Возможно на ендпоинт сделать одну функцию с разделением на методы внутри функции.
	router := mux.NewRouter()

	router.HandleFunc("/api/books", getAllBooks).Methods("GET")
	// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	serv := http.Server{
		Addr:    net.JoinHostPort("", os.Getenv("PORT")),
		Handler: router,
	}
	serv.ListenAndServe()
}
