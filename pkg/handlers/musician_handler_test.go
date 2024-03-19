package handlers

import (
	"bytes"
	"encoding/json"
	"models"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestUpdateMusician(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a sample musician
	existingMusician := models.Musician{Name: "John Doe", MusicianType: "Guitarist"}

	// Mock the BindJSON function to decode the JSON payload into a Musician struct
	gin.SetMode(gin.TestMode)
	router.PUT("/musicians/:musicianName", func(c *gin.Context) {
		// Simulate the musician name from the URL param
		musicianName := c.Param("musicianName")

		// Check if the musician exists
		if musicianName != existingMusician.Name {
			c.JSON(http.StatusNotFound, gin.H{"error": "Musician not found"})
			return
		}

		// Decode the JSON payload into a Musician struct
		var updatedMusician models.Musician
		if err := c.BindJSON(&updatedMusician); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update the musician details
		existingMusician = updatedMusician

		// Respond with the updated musician details
		c.JSON(http.StatusOK, gin.H{"message": "Musician updated successfully", "musician": updatedMusician})
	})

	// Create a sample JSON payload representing the updated musician information
	updatedMusician := models.Musician{Name: "John Doe", MusicianType: "Pianist"}
	payload, _ := json.Marshal(updatedMusician)

	// Create a new HTTP PUT request with the JSON payload
	req, err := http.NewRequest("PUT", "/musicians/"+existingMusician.Name, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request and record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the response status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Decode the response body into a struct to check the updated musician details
	var response struct {
		Message  string          `json:"message"`
		Musician models.Musician `json:"musician"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the response message and updated musician details
	assert.Equal(t, "Musician updated successfully", response.Message)
	assert.Equal(t, updatedMusician, response.Musician)
}

func TestGetAlbumsForMusicianSortedByPrice(t *testing.T) {
	// Define a sample musician name
	musicianName := "MusicianName"

	// Define a sample list of albums
	// testAlbums := []models.MusicAlbum{
	// 	{AlbumName: "Album 1", Price: 10, Musicians: []string{musicianName}},
	// 	{AlbumName: "Album 2", Price: 5, Musicians: []string{"AnotherMusician"}},
	// 	{AlbumName: "Album 3", Price: 15, Musicians: []string{musicianName}},
	// }

	// Create a new Gin router
	router := gin.Default()

	// Define a mock endpoint to test the GetAlbumsForMusicianSortedByPrice function
	gin.SetMode(gin.TestMode)
	router.GET("/albumsForMusicianSortedByPrice/:musicianName", func(c *gin.Context) {
		// Call the GetAlbumsForMusicianSortedByPrice function with the provided musician name
		GetAlbumsForMusicianSortedByPrice(c)
	})

	// Create a new HTTP GET request to test getting albums for a musician sorted by price
	req, err := http.NewRequest("GET", "/albumsForMusicianSortedByPrice/"+musicianName, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request and record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the response status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Decode the response body into a slice of albums
	var responseAlbums []models.MusicAlbum
	err = json.Unmarshal(rec.Body.Bytes(), &responseAlbums)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the returned albums are sorted by price in ascending order
	assert.True(t, sort.SliceIsSorted(responseAlbums, func(i, j int) bool {
		return responseAlbums[i].Price < responseAlbums[j].Price
	}))

	// Assert that all returned albums belong to the specified musician
	for _, album := range responseAlbums {
		assert.Contains(t, album.Musicians, musicianName)
	}
}

func setupSampleAlbums() []models.MusicAlbum {
	// Define sample albums with musicians
	albums := []models.MusicAlbum{
		{
			AlbumName:     "Album1",
			DateOfRelease: time.Now(),
			Genre:         "Genre1",
			Price:         10,
			Description:   "Description1",
			Musicians:     []string{"Alice", "Bob", "Charlie"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album3",
			DateOfRelease: time.Now(),
			Genre:         "Genre3",
			Price:         30,
			Description:   "Description3",
			Musicians:     []string{"David", "Eve", "Frank", "Sinatra"},
		},
		{
			AlbumName:     "Album4",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         60,
			Description:   "Description2",
			Musicians:     []string{"Ross", "Joey", "Frank"},
		},
		{
			AlbumName:     "Album5",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         100,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
		{
			AlbumName:     "Album2",
			DateOfRelease: time.Now(),
			Genre:         "Genre2",
			Price:         20,
			Description:   "Description2",
			Musicians:     []string{"David", "Eve", "Frank"},
		},
	}

	return albums
}

func TestGetMusiciansForAlbumSortedByName(t *testing.T) {
	// Populate sample albums with musicians
	albums := setupSampleAlbums()

	// Setup Gin
	r := gin.Default()

	// Define the test handler
	r.GET("/musicians/:albumName", func(c *gin.Context) {
		// Get the album name from the request parameters
		albumName := c.Param("albumName")

		// Find the album with the specified name
		var album models.MusicAlbum
		for _, a := range albums {
			if a.AlbumName == albumName {
				album = a
				break
			}
		}

		// If album not found, return empty musicians list
		if album.AlbumName == "" {
			c.JSON(http.StatusOK, []string{})
			return
		}

		// Sort the musicians by name in ascending order
		sort.Strings(album.Musicians)

		// Return sorted musicians list
		c.JSON(http.StatusOK, album.Musicians)
	})

	// Define test data
	albumName := "Album1"
	expectedMusicians := []string{"Alice", "Bob", "Charlie"}

	// Define the request
	req, err := http.NewRequest("GET", "/musicians/"+albumName, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the HTTP response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body
	var response []string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Sort the expected musicians
	sort.Strings(expectedMusicians)

	// Check if the response is correct
	assert.Equal(t, expectedMusicians, response)
}
