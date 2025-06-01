package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func getPost(c *gin.Context) {
	id := c.Param("id")

	for _, p := range posts {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Welcome to GoCharity 🍷",
		})
	})

	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)

	log.Printf("🚀 Server is running at http://localhost:%s", port)
	router.Run(":" + port)
}
