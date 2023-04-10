package controllers

import (
	"Assignment3/Init"
	"Assignment3/models"
	"github.com/gin-gonic/gin"
)

func BookCreate(c *gin.Context) {
	var body struct {
		ID          uint
		Title       string
		Description string
		Cost        uint
	}
	c.Bind(&body)

	post := models.Book{ID: body.ID, Title: body.Title, Description: body.Description, Cost: body.Cost}
	Init.DB.Create(&post)
	c.JSON(200, gin.H{
		"book": post,
	})
}
func BooksId(c *gin.Context) {

	id := c.Param("id")
	var books models.Book
	Init.DB.First(&books, id)

	c.JSON(200, gin.H{
		"books": books,
	})
}
func BooksShow(c *gin.Context) {

	var books []models.Book
	Init.DB.Find(&books)

	c.JSON(200, gin.H{
		"books": books,
	})
}
func BooksUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		ID          uint
		Title       string
		Description string
		Cost        uint
	}
	c.Bind(&body)

	var book models.Book
	Init.DB.First(&book, id)

	Init.DB.Model(&book).Updates(models.Book{
		ID:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Cost:        body.Cost,
	})

	c.JSON(200, gin.H{
		"books": book,
	})
}
func BooksDelete(c *gin.Context) {

	id := c.Param("id")
	Init.DB.Delete(&models.Book{}, id)

	c.Status(200)
}
func BooksSortAsc(c *gin.Context) {
	var books []models.Book
	Init.DB.Order("cost asc").Find(&books)

	c.JSON(200, gin.H{
		"books": books,
	})
}
func BooksSortDesc(c *gin.Context) {
	var books []models.Book
	Init.DB.Order("cost desc").Find(&books)

	c.JSON(200, gin.H{
		"books": books,
	})
}
