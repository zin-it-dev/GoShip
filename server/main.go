package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Post struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
}

var posts []Post

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

	data, err := os.ReadFile("./data/posts.json")
	if err != nil {
		log.Fatalf("Error reading posts.json: %v", err)
	}

	err = json.Unmarshal(data, &posts)
	if err != nil {
		log.Fatalf("Error parsing posts.json: %v", err)
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