package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("%+v\n", albums)
	
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = [] Album{	
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.90}, 
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan",Price:23.43},
}

