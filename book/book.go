package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/sidecut/go-fiber-app/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
