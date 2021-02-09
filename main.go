package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sidecut/go-fiber-app/book"
	"github.com/sidecut/go-fiber-app/database"
)

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func setupRoutes(app *gin.Engine) {
	app.GET("/", helloWorld)

	app.GET("/api/v1/book", book.GetBooks)
	app.GET("/api/v1/book/:id", book.GetBook)
	app.POST("/api/v1/book", book.NewBook)
	app.DELETE("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	r := gin.Default()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(r)
	r.Run()
}
