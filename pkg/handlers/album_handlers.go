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
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAlbum.DateOfRelease = time.Now()
	albums = append(albums, newAlbum)
	response := struct {
		Message string            `json:"message"`
		Album   models.MusicAlbum `json:"MusicAlbum"`
	}{
		Message: "Album created successfully",
		Album:   newAlbum,
	}

	c.IndentedJSON(http.StatusCreated, response.Message)
}

func UpdateAlbum(c *gin.Context) {
	albumIndex := FindAlbumIndex(c.Param("albumName"))
	if albumIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	originalAlbum := albums[albumIndex]

	var updatedAlbum models.MusicAlbum
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	updatedAlbum.DateOfRelease = originalAlbum.DateOfRelease

	albums[albumIndex] = updatedAlbum

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
