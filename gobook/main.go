package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	client := ConnectDB()

	router.LoadHTMLGlob("views/*")
	router.Static("/public", "./public")

	router.GET("/books", func(c *gin.Context) {
		books, err := handleGetBooks(c, client)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to query books"})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"books": books,
		})
	})
	router.GET("/books/:uuid", func(c *gin.Context) {
		handleGetById(c, client)
	})
	router.POST("/books", func(c *gin.Context) {
		handlePostBooks(c, client)
	})
	router.DELETE("/books/:uuid", func(c *gin.Context) {
		handleDeleteBook(c, client)
	})

	router.Run(":8080")
}
