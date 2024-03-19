package handlers

import (
	"models"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

var albums = []models.MusicAlbum{}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.MusicAlbum
	newAlbum.DateOfRelease = time.Now()
	// Bind JSON data to the newAlbum variable
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate album name
	if len(newAlbum.AlbumName) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album name must be at least 5 characters"})
		return
	}

	// Validate date of release (mandatory)
	if newAlbum.DateOfRelease.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date of release is mandatory"})
		return
	}

	// Validate genre (optional)

	// Validate price
	if newAlbum.Price < 100 || newAlbum.Price > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be between 100 and 1000"})
		return
	}

	// Validate description (optional)

	// If all validations pass, add the album to the albums slice

	albums = append(albums, newAlbum)

	// Prepare response
	response := struct {
		Message string            `json:"message"`
		Album   models.MusicAlbum `json:"MusicAlbum"`
	}{
		Message: "Album created successfully",
		Album:   newAlbum,
	}

	c.IndentedJSON(http.StatusCreated, response.Album)
}

func UpdateAlbum(c *gin.Context) {
	// Find the index of the album to be updated
	albumIndex := FindAlbumIndex(c.Param("albumName"))
	if albumIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Get the original album from the albums slice
	originalAlbum := albums[albumIndex]

	// Bind the updated album data from the request
	var updatedAlbum models.MusicAlbum
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate updated album name
	if len(updatedAlbum.AlbumName) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album name must be at least 5 characters"})
		return
	}

	// Validate updated price
	if updatedAlbum.Price < 100 || updatedAlbum.Price > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be between 100 and 1000"})
		return
	}

	// Validate updated description (optional)

	// Set the updated date of release to match the original album
	updatedAlbum.DateOfRelease = originalAlbum.DateOfRelease

	// Update the album in the albums slice
	albums[albumIndex] = updatedAlbum

	// Return success response with the updated album
	c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully", "album": updatedAlbum})
}

func FindAlbumIndex(albumName string) int {
	for i, album := range albums {
		if album.AlbumName == albumName {
			return i
		}
	}
	return -1
}

func GetAlbumsSortedByDate(c *gin.Context) {
	// Create a copy of the albums slice to avoid modifying the original slice
	sortedAlbums := make([]models.MusicAlbum, len(albums))
	copy(sortedAlbums, albums)

	// Sort the albums by DateOfRelease in ascending order
	sort.Slice(sortedAlbums, func(i, j int) bool {
		return sortedAlbums[i].DateOfRelease.Before(sortedAlbums[j].DateOfRelease)
	})

	c.IndentedJSON(http.StatusOK, sortedAlbums)
}
