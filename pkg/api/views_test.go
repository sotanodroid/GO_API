package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sotanodroid/GO_API/pkg/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	models.InitDB("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
}

func TestGetRequest(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()

	getAllBooks(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestCreateBook(t *testing.T) {
	book := models.Book{
		Isbn:  "123459",
		Title: "test Book",
		Author: models.Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	}

	requestByte, _ := json.Marshal(book)

	r, _ := http.NewRequest("POST", "/api/books", bytes.NewReader(requestByte))
	w := httptest.NewRecorder()

	createBook(w, r)

	respBook := new(models.Book)
	_ = json.NewDecoder(w.Body).Decode(&respBook)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, book, *respBook)
}
