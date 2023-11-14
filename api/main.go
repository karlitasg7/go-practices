package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    string `json: "id"`
	Title string `json: "title"`
}

var books = []book{
	{ID: "1", Title: "Book1"},
	{ID: "3", Title: "Book2"},
	{ID: "4", Title: "Book3"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBooksById(c *gin.Context) {

	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func postBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBooksById)
	router.POST("/books", postBook)

	router.Run("localhost:8080")

}
