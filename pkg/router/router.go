package router

import (
	"handlers"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()

	// Musicians Endpoints
	musiciansGroup := router.Group("/musicians")
	{
		musiciansGroup.GET("", handlers.GetMusicians)
		musiciansGroup.POST("", handlers.CreateMusician)
		musiciansGroup.PATCH("/:musicianName", handlers.UpdateMusician)
	}

	// Albums Endpoints
	albumsGroup := router.Group("/albums")
	{
		albumsGroup.GET("", handlers.GetAlbums)
		albumsGroup.POST("", handlers.CreateAlbum)
		albumsGroup.PATCH("/:albumName", handlers.UpdateAlbum)
		albumsGroup.GET("/sortedByDate", handlers.GetAlbumsSortedByDate)
		albumsGroup.GET("/albumsForMusicianSortedByPrice/:musicianName", handlers.GetAlbumsForMusicianSortedByPrice)
		albumsGroup.GET("/musicians/:albumName", handlers.GetMusiciansForAlbumSortedByName)
	}

	router.Run(":8080")
}
