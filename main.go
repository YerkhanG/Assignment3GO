package main

import (
	"Assignment3/Init"
	"Assignment3/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	Init.LoadEnvVar()
	Init.ConnectDB()
}
func main() {
	r := gin.Default()
	r.PUT("/books/:id", controllers.BooksUpdate)
	r.GET("/books", controllers.BooksShow)
	r.GET("/books/asc", controllers.BooksSortAsc)
	r.GET("/books/desc", controllers.BooksSortDesc)
	r.POST("/books", controllers.BookCreate)
	r.GET("/books/:id", controllers.BooksId)
	r.DELETE("/books/:id", controllers.BooksDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
