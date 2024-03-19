package handlers

import (
	"models"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

var musicians = []models.Musician{}

func GetMusicians(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, musicians)
}

func CreateMusician(c *gin.Context) {
	var newMusician models.Musician
	if err := c.BindJSON(&newMusician); err != nil {
		return
	}

	if len(newMusician.Name) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Musician name must be at least 3 characters"})
		return
	}

	musicians = append(musicians, newMusician)
	response := struct {
		Message  string          `json:"message"`
		Musician models.Musician `json:"musician"`
	}{
		Message:  "Musician created successfully",
		Musician: newMusician,
	}

	c.IndentedJSON(http.StatusCreated, response)
}

func UpdateMusician(c *gin.Context) {
	musicianIndex := FindMusicianIndex(c.Param("musicianName"))
	if musicianIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Musician not found"})
		return
	}

	var updatedMusician models.Musician
	if err := c.BindJSON(&updatedMusician); err != nil {
		return
	}

	musicians[musicianIndex] = updatedMusician
	c.JSON(http.StatusOK, gin.H{"message": "Musician updated successfully", "musician": updatedMusician})
}

func FindMusicianIndex(musicianName string) int {
	for i, musician := range musicians {
		if musician.Name == musicianName {
			return i
		}
	}
	return -1
}

func GetAlbumsForMusicianSortedByPrice(c *gin.Context) {
	musicianName := c.Param("musicianName")
	var musicianAlbums []models.MusicAlbum

	// Iterate through albums to find albums for the specified musician
	for _, album := range albums {
		for _, musician := range album.Musicians {
			if musician == musicianName {
				musicianAlbums = append(musicianAlbums, album)
				break
			}
		}
	}

	// Sort the musician's albums by price in ascending order
	sort.Slice(musicianAlbums, func(i, j int) bool {
		return musicianAlbums[i].Price < musicianAlbums[j].Price
	})

	c.IndentedJSON(http.StatusOK, musicianAlbums)
}

func GetMusiciansForAlbumSortedByName(c *gin.Context) {
	albumName := c.Param("albumName")
	var albumMusicians []string

	// Find the album with the specified name
	for _, album := range albums {
		if album.AlbumName == albumName {
			albumMusicians = album.Musicians
			break
		}
	}

	// Sort the musicians by name in ascending order
	sort.Strings(albumMusicians)

	c.IndentedJSON(http.StatusOK, albumMusicians)
}
