package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumById)
    router.POST("/albums", postAlbums)

    router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Try to bind the body of the request to an album
	if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, current_album := range albums {
		if current_album.ID == id {
			c.IndentedJSON(http.StatusOK, current_album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
