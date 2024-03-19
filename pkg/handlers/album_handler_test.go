package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"models"
)

func TestCreateAlbum(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Set up a mock request
	requestBody := models.MusicAlbum{
		AlbumName:   "4",
		Genre:       "William pearce",
		Price:       2,
		Description: "nce",
		Musicians:   []string{"Josn", "Alice"},
	}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Set up a mock response recorder
	w := httptest.NewRecorder()

	// Mock the CreateAlbum handler function
	router.POST("/albums", CreateAlbum)

	// Dispatch the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the response body
	var response models.MusicAlbum
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Check the created album details
	assert.Equal(t, requestBody.AlbumName, response.AlbumName)
	assert.Equal(t, requestBody.Genre, response.Genre)
	assert.Equal(t, requestBody.Price, response.Price)
	assert.Equal(t, requestBody.Description, response.Description)
	assert.Equal(t, requestBody.Musicians, response.Musicians)

	// Check that the DateOfRelease is not zero
	assert.NotEqual(t, time.Time{}, response.DateOfRelease)
}

func TestUpdateAlbum(t *testing.T) {
	// Set up a test router using Gin
	router := gin.Default()

	// Define a route to create an album
	router.POST("/albums", CreateAlbum)

	// Define a route to update an album
	router.PUT("/albums/:albumName", UpdateAlbum)

	// Create a request to add an album with name "4"
	createAlbumRequest := httptest.NewRequest("POST", "/albums", bytes.NewBufferString(`{
		"albumName": "4",
		"genre": "William pearce",
		"price": 2,
		"description": "nce",
		"musicians": ["Josn", "Alice"]
	}`))
	createAlbumResponse := httptest.NewRecorder()
	router.ServeHTTP(createAlbumResponse, createAlbumRequest)

	// Assert that the album was created successfully (status code 201)
	assert.Equal(t, http.StatusCreated, createAlbumResponse.Code)

	// Create a request to update the album with name "4"
	updateAlbumRequest := httptest.NewRequest("PUT", "/albums/4", bytes.NewBufferString(`{
		"albumName": "4",
		"genre": "Updated Genre",
		"price": 5,
		"description": "Updated Description",
		"musicians": ["Updated Musician"]
	}`))
	updateAlbumResponse := httptest.NewRecorder()
	router.ServeHTTP(updateAlbumResponse, updateAlbumRequest)

	// Assert that the album was updated successfully (status code 200)
	assert.Equal(t, http.StatusOK, updateAlbumResponse.Code)
}

func TestGetAlbumsSortedByDate(t *testing.T) {
	// Mock data: create unsorted albums
	unsortedAlbums := []models.MusicAlbum{
		{AlbumName: "Album3", DateOfRelease: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		{AlbumName: "Album1", DateOfRelease: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		{AlbumName: "Album2", DateOfRelease: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// Create a Gin router
	router := gin.Default()

	// Set up the route handler
	router.GET("/getAlbumsSortedByDate", GetAlbumsSortedByDate)

	// Create a request to the handler
	req, err := http.NewRequest("GET", "/getAlbumsSortedByDate", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request to the handler
	router.ServeHTTP(w, req)

	// Assert that the response status code is OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Unmarshal the response body into sortedAlbums
	var sortedAlbums []models.MusicAlbum
	err = json.Unmarshal(w.Body.Bytes(), &sortedAlbums)
	if err != nil {
		t.Fatal(err)
	}

	// Sort the unsortedAlbums slice to compare with sortedAlbums
	sort.Slice(unsortedAlbums, func(i, j int) bool {
		return unsortedAlbums[i].DateOfRelease.Before(unsortedAlbums[j].DateOfRelease)
	})

	// Assert that the sortedAlbums and unsortedAlbums are equal
	for i := 0; i < len(sortedAlbums)-1; i++ {
		assert.True(t, sortedAlbums[i].DateOfRelease.Before(sortedAlbums[i+1].DateOfRelease))
	}
}
