package book

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sidecut/go-fiber-app/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *gin.Context) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(http.StatusOK, book)
}

func NewBook(c *gin.Context) {
	db := database.DBConn
	book := new(Book)
	if err := c.BindJSON(book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.AbortWithError(http.StatusNotFound, errors.New("No Book Found with ID")) //,  "No Book Found with ID"))
		return
	}
	db.Delete(&book)
	c.String(http.StatusOK, "Book Successfully deleted")
}
