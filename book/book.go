package book

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sidecut/go-fiber-app/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c echo.Context) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.QueryParams()["id"]
	db := database.DBConn
	var book Book
	result := db.Find(&book, id)
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return c.String(http.StatusNotFound, result.Error.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func NewBook(c echo.Context) error {
	db := database.DBConn
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(err.(*echo.HTTPError).Code, err)
	}
	db.Create(&book)
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) (err error) {
	id := c.QueryParams()["id"]
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.String(http.StatusNotFound, "No Book Found with ID")
		return
	}
	db.Delete(&book)
	return c.String(http.StatusOK, "Book Successfully deleted")
}
