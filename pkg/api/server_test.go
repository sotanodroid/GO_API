package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/sotanodroid/GO_API/pkg/models"
	"github.com/stretchr/testify/assert"
)

type BookSlice struct {
	Books []models.Book
}

func TestServer(t *testing.T) {
	srv, ctx := setup()
	endpoints := MakeEndpoints(srv)
	handler := NewHTTPServer(ctx, endpoints)
	allBooks := &BookSlice{}

	{
		payload := strings.NewReader(`{
			"isbn":"4545454",
			"title":"Book Three",
			"author": {
				"firstname":"John",
				"lastname":"Doe"
			  }
		  }`)
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/api/books", payload)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "{\"ok\":\"Created\"}\n", recorder.Body.String())
	}

	{
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/api/books", nil)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), allBooks)
	}

	lastBook := allBooks.Books[len(allBooks.Books)-1]

	{
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/api/books/"+strconv.Itoa(lastBook.ID), nil)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		book := new(models.Book)
		json.Unmarshal(recorder.Body.Bytes(), book)

		assert.Equal(t, lastBook.ID, book.ID)
		assert.Equal(t, lastBook.Isbn, book.Isbn)
		assert.Equal(t, lastBook.Title, book.Title)
		assert.Equal(t, lastBook.Author, book.Author)
	}
	{
		payload := strings.NewReader(`{
			"isbn":"111111",
			"title":"Updated"
		}`)

		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/api/books/"+strconv.Itoa(lastBook.ID), payload)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		assert.Equal(t, "{\"ok\":\"Updated\"}\n", recorder.Body.String())
	}
	{
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/api/books/"+strconv.Itoa(lastBook.ID), nil)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		assert.Equal(t, "{\"ok\":\"Deleted\"}\n", recorder.Body.String())
	}
}
