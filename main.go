package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sidecut/go-fiber-app/book"
	"github.com/sidecut/go-fiber-app/database"
)

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func setupRoutes(app *echo.Echo) {
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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
