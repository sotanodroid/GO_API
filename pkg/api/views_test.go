package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sotanodroid/GO_API/pkg/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	models.InitDB("postgres://postgres:postgres@localhost:5432/postgres")
}

func TestGetRequest(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()

	getAllBooks(w, r)

	// TODO assert output
	assert.Equal(t, http.StatusOK, w.Code)
}

// TODO Test for all endpoints
