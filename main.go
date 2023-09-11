package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int
	Name   string
	Author string
}

var books []Book

func main() {
	req := gin.Default()
	req.GET("/books", getBooks)
	req.GET("/book/:id", getBook)
	req.POST("/book", insertBook)
	req.POST("/books", insertBooks)
	req.PUT("/book", editBook)
	req.DELETE("/book/:id", deleteBook)
	req.Run(":3000")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}
func getBook(c *gin.Context) {
	var isFound bool
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	for _, book := range books {
		if book.ID == ID {
			isFound = true
			c.JSON(http.StatusOK, book)
			break
		}
	}
	if !isFound {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "not found book",
		})
	}

}

func insertBook(c *gin.Context) {
	var book Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	var max int
	for _, v := range books {
		if max < v.ID {
			max = v.ID
		}
	}
	max++
	book.ID = max + 1
	books = append(books, book)
	c.JSON(http.StatusOK, book)
}

func insertBooks(c *gin.Context) {
	var bookReqs []Book
	err := c.ShouldBind(&bookReqs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	var max int
	for _, v := range books {
		if max < v.ID {
			max = v.ID
		}
	}
	max++
	for i, _ := range bookReqs {
		bookReqs[i].ID = max
		max++
	}
	books = append(books, bookReqs...)
	c.JSON(http.StatusOK, bookReqs)
}

func editBook(c *gin.Context) {
	var bookReq Book
	var isFound bool
	err := c.ShouldBind(&bookReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	for i, book := range books {
		if book.ID == bookReq.ID {
			books[i].Author, books[i].Name = bookReq.Author, bookReq.Name
			c.JSON(http.StatusOK, bookReq)
			isFound = true
			break
		}
	}
	if !isFound {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "not found book",
		})
	}
}

func deleteBook(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	var isFound bool
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	for i, book := range books {
		if book.ID == int(ID) {
			isFound = true
			books = removeBook(books, i)
			c.JSON(http.StatusOK, ID)
			break
		}
	}
	if !isFound {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "not found book",
		})
	}

}

func removeBook(books []Book, index int) []Book {
	return append(books[:index], books[index+1:]...)
}
