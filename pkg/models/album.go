package models

import "time"

type MusicAlbum struct {
	AlbumName     string    `json:"albumName"`
	DateOfRelease time.Time `json:"releaseDate"`
	Genre         string    `json:"genre"`
	Price         int       `json:"price"`
	Description   string    `json:"description"`
	Musicians     []string  `json:"musicians"`
}
