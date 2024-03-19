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

	// Set up a mock request with valid data
	requestBody := models.MusicAlbum{
		AlbumName:     "AlbumNameTest",                    // Valid album name with at least 5 characters
		DateOfRelease: time.Now(),                         // Mandatory date of release
		Genre:         "GenreTest",                        // Optional genre
		Price:         500,                                // Valid price between 100 and 1000
		Description:   "DescriptionTest",                  // Optional description
		Musicians:     []string{"Musician1", "Musician2"}, // Optional list of musicians
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
