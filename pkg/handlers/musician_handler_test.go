package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	// Update with the correct package name
	// Update with the correct package name
)

func TestCreateMusician(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for testing CreateMusician function
	router.POST("/musicians", CreateMusician)

	// Prepare the request body with JSON data
	jsonData := []byte(`{
		"name": "John Doe",
		"musicianType": "Guitarist"
	}`)

	// Create an HTTP request with the prepared JSON data
	req, err := http.NewRequest("POST", "/musicians", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request handler function
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	expectedResponse := `{"message":"Musician created successfully","musician":{"name":"John Doe","musicianType":"Guitarist"}}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
