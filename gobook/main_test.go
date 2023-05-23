package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRuter() (*gin.Engine, *mongo.Client) {
	router := gin.Default()
	client := ConnectDB()
	return router, client
}

func TestPostBookHandler(t *testing.T) {
	router, client := SetUpRuter()
	router.POST("/books", func(c *gin.Context) {
		handlePostBooks(c, client)
	})
	book := NewBook("Test", "Test Author", time.Now(), "Test Publisher")
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBooksHandler(t *testing.T) {
	router, client := SetUpRuter()
	router.GET("/books", func(c *gin.Context) {
		handleGetBooks(c, client)
	})

	req, _ := http.NewRequest("GET", "/books", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBooksHandler(t *testing.T) {
	router, client := SetUpRuter()
	router.POST("/books", func(c *gin.Context) {
		handlePostBooks(c, client)
	})
	router.DELETE("/books/:uuid", func(c *gin.Context) {
		handleDeleteBook(c, client)
	})

	book := NewBook("Test", "Test Author", time.Now(), "Test Publisher")
	jsonValue, _ := json.Marshal(book)

	// Create a new book
	reqCreate, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	wCreate := httptest.NewRecorder()
	router.ServeHTTP(wCreate, reqCreate)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	// Unmarshal the response body to get the ID of the created book
	var createdBook Book
	json.Unmarshal(wCreate.Body.Bytes(), &createdBook)

	// Delete the created book
	reqDel, _ := http.NewRequest("DELETE", "/books/"+createdBook.ID.String(), bytes.NewBuffer(jsonValue))
	wDel := httptest.NewRecorder()
	router.ServeHTTP(wDel, reqDel)

	assert.Equal(t, http.StatusOK, wDel.Code)
}
