package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
}

var posts = []Post{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane"},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan"},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan"},
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Welcome to GoCharity 🍷",
		})
	})

	router.GET("/posts", getPosts)

	router.Run("localhost:8080")
}