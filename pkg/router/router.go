package router

import (
	"handlers"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	router.GET("/getAlbums", handlers.GetAlbums)
	router.POST("/createAlbum", handlers.CreateAlbum)
	router.PUT("/updateAlbum/:albumName", handlers.UpdateAlbum)
	router.GET("/getAlbumsSortedByDate", handlers.GetAlbumsSortedByDate)
	router.GET("/getMusicians", handlers.GetMusicians)
	router.POST("/createMusician", handlers.CreateMusician)
	router.PUT("/updateMusician/:musicianName", handlers.UpdateMusician)
	router.GET("/getAlbumsForMusicianSortedByPrice/:musicianName", handlers.GetAlbumsForMusicianSortedByPrice)
	router.GET("/getMusiciansForAlbumSortedByName/:albumName", handlers.GetMusiciansForAlbumSortedByName)
	router.Run("localhost:8080")
}
